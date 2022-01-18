package main

import (
	"context"
	"go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

/**
测试服务注册一次
*/

// 创建租约注册服务
type ServiceOnceReg struct {
	client     *clientv3.Client
	lease      clientv3.Lease
	leaseResp  *clientv3.LeaseGrantResponse
	canclefunc func()
	key        string
}

// 新建 服务注册
func NewServiceOnceReg(addr []string, timeNum int64) (*ServiceOnceReg, error) {
	conf := clientv3.Config{
		Endpoints:   addr,
		DialTimeout: time.Duration(5) * time.Second,
	}

	var (
		client *clientv3.Client
	)

	if clientTem, errCN := clientv3.New(conf); errCN == nil {
		client = clientTem
	} else {
		return nil, errCN
	}

	ser := &ServiceOnceReg{
		client: client,
	}

	if errSS := ser.setLease(timeNum); errSS != nil {
		return nil, errSS
	}
	return ser, nil
}

// 设置租约
func (sr *ServiceOnceReg) setLease(timeNum int64) error {

	lease := clientv3.NewLease(sr.client)

	// 设置租约时间
	leaseResp, errLG := lease.Grant(context.TODO(), timeNum)
	if errLG != nil {
		return errLG
	}

	// 设置续租
	_, cancelFunc := context.WithCancel(context.TODO())

	sr.lease = lease
	sr.leaseResp = leaseResp
	sr.canclefunc = cancelFunc
	return nil
}

// 设置租约
func (sr *ServiceOnceReg) setLeaseOnce() error {

	// 设置续租
	ctx, cancelFunc := context.WithCancel(context.TODO())

	// 自动续约时间一次，从KeepAliveOnce 函数被调用前开始
	log.Println("Start lease once.")
	leaseRespChan, errLK := sr.lease.KeepAliveOnce(ctx, sr.leaseResp.ID)
	log.Println(leaseRespChan)

	if errLK != nil {
		return errLK
	}

	sr.canclefunc = cancelFunc
	return nil
}

// 通过租约 注册服务
func (sr *ServiceOnceReg) PutService(key, val string) error {
	kv := clientv3.NewKV(sr.client)
	_, errKP := kv.Put(context.TODO(), key, val, clientv3.WithLease(sr.leaseResp.ID))
	return errKP
}

// 通过租约 删除服务
func (sr *ServiceOnceReg) DeleteService(key string) error {
	kv := clientv3.NewKV(sr.client)
	_, errKP := kv.Delete(context.TODO(), key)
	return errKP
}

// 撤销租约
func (sr *ServiceOnceReg) RevokeLease() error {

	log.Printf("[%d] start revoke service.", sr.leaseResp.ID)

	sr.canclefunc()
	time.Sleep(time.Duration(2) * time.Second)
	_, errSLR := sr.lease.Revoke(context.TODO(), sr.leaseResp.ID)
	return errSLR
}

func main() {
	ser, _ := NewServiceOnceReg([]string{"10.122.105.131:12379", "10.122.105.131:22379", "10.122.105.131:32379"}, 5)
	ser.PutService("/node/999", "kakaka")

	time.Sleep(time.Duration(3) * time.Second)

	// 需要在TTL 结束前设置，否则etcdserver: requested lease not found.
	errSS := ser.setLeaseOnce()
	if errSS != nil {
		log.Printf("Set lease once err: %v.", errSS)
	}

	select {}

}
