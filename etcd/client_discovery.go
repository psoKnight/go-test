package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	"go.etcd.io/etcd/client/v3"
	"log"
	"sync"
	"time"
)

type ClientDis struct {
	client     *clientv3.Client
	serverList map[string]string // 当前的注册服务
	lock       sync.Mutex
}

// NewClientDis 新建 服务发现
func NewClientDis(addr []string) (*ClientDis, error) {
	conf := clientv3.Config{
		Endpoints:   addr,
		DialTimeout: 5 * time.Second,
	}
	if client, errCN := clientv3.New(conf); errCN == nil {
		return &ClientDis{
			client:     client,
			serverList: make(map[string]string),
		}, nil
	} else {
		return nil, errCN
	}
}

// GetService 获取服务
func (sr *ClientDis) GetService(prefix string) ([]string, error) {

	// 获取键，返回对应的值，此处为只返回匹配指定前缀的值
	// 获取当前
	resp, errSG := sr.client.Get(context.Background(), prefix, clientv3.WithPrefix())
	if errSG != nil {
		return nil, errSG
	}

	// 获取值
	// 获取动态增长
	addrs := sr.extractAddrs(resp)

	go sr.watcher(prefix)
	return addrs, nil
}

// 监听
func (sr *ClientDis) watcher(prefix string) {

	// 监听键，返回对应的值，此处为只返回匹配指定前缀的值
	rch := sr.client.Watch(context.Background(), prefix, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case mvccpb.PUT:
				sr.SetServiceList(string(ev.Kv.Key), string(ev.Kv.Value))
			case mvccpb.DELETE:
				sr.DelServiceList(string(ev.Kv.Key))
			}
		}
	}
}

// 转换值的存储格式
func (sr *ClientDis) extractAddrs(resp *clientv3.GetResponse) []string {
	addrs := make([]string, 0)
	if resp == nil || resp.Kvs == nil {
		return addrs
	}
	for i := range resp.Kvs {
		if v := resp.Kvs[i].Value; v != nil {
			sr.SetServiceList(string(resp.Kvs[i].Key), string(resp.Kvs[i].Value))
			addrs = append(addrs, string(v))
		}
	}
	return addrs
}

// SetServiceList 存储当前的注册服务
func (sr *ClientDis) SetServiceList(key, val string) {
	sr.lock.Lock()
	defer sr.lock.Unlock()
	sr.serverList[key] = string(val)
	log.Println(fmt.Sprintf("Set data key: %s, val: %s.", key, val))
}

// DelServiceList 删除当前的存储服务
func (sr *ClientDis) DelServiceList(key string) {
	sr.lock.Lock()
	defer sr.lock.Unlock()
	delete(sr.serverList, key)
	log.Println(fmt.Sprintf("Del data key: %s.", key))
}

// SerList2Array 返回存储服务
func (sr *ClientDis) SerList2Array() []string {
	sr.lock.Lock()
	defer sr.lock.Unlock()
	addrs := make([]string, 0)

	for _, v := range sr.serverList {
		addrs = append(addrs, v)
	}
	return addrs
}

//Close 关闭服务
func (sr *ClientDis) Close() error {
	return sr.client.Close()
}

func main() {
	cli, _ := NewClientDis([]string{"10.117.49.69:12379", "10.117.49.69:22379", "10.117.49.69:32379"})
	cli.GetService("/node")

	count := 0
	for {
		log.Println(fmt.Sprintf("Current list: %v.", cli.SerList2Array()))

		time.Sleep(time.Duration(1) * time.Second)

		count++
		if count == 10 {
			log.Println("Close discover.")
			cli.Close()
			return
		}
	}
}
