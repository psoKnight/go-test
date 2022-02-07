package main

import (
	"context"
	"go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

// 创建租约注册服务
type ServiceReg struct {
	client        *clientv3.Client
	lease         clientv3.Lease
	leaseResp     *clientv3.LeaseGrantResponse
	canclefunc    func()
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
	key           string
}

// 新建 服务注册
func NewServiceReg(addr []string, timeNum int64) (*ServiceReg, error) {
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

	ser := &ServiceReg{
		client: client,
	}

	if errSS := ser.setLease(timeNum); errSS != nil {
		return nil, errSS
	}
	go ser.ListenLeaseRespChan()
	return ser, nil
}

// 设置租约
func (sr *ServiceReg) setLease(timeNum int64) error {
	lease := clientv3.NewLease(sr.client)

	// 设置租约时间
	leaseResp, errLG := lease.Grant(context.TODO(), timeNum)
	if errLG != nil {
		return errLG
	}

	// 设置续租
	ctx, cancelFunc := context.WithCancel(context.TODO())

	// 自动续约时间为：TTL 时间/3
	leaseRespChan, errLK := lease.KeepAlive(ctx, leaseResp.ID)

	if errLK != nil {
		return errLK
	}

	sr.lease = lease
	sr.leaseResp = leaseResp
	sr.canclefunc = cancelFunc
	sr.keepAliveChan = leaseRespChan
	return nil
}

// 监听 续租情况
func (sr *ServiceReg) ListenLeaseRespChan() {
	for {
		select {

		// 时间间隔为收到chan 信息的时刻
		case leaseKeepResp := <-sr.keepAliveChan:
			if leaseKeepResp == nil {
				log.Printf("'%d'Lease renewal has been closed.\n", sr.leaseResp.ID)
				return
			} else {
				log.Printf("Lease '%d' success, detail: %+v.\n", sr.leaseResp.ID, leaseKeepResp)

			}
		}
	}
}

// 通过租约 注册服务
func (sr *ServiceReg) PutService(key, val string) error {
	kv := clientv3.NewKV(sr.client)
	_, errKP := kv.Put(context.TODO(), key, val, clientv3.WithLease(sr.leaseResp.ID))
	return errKP
}

// 通过租约 删除服务
func (sr *ServiceReg) DeleteService(key string) error {
	kv := clientv3.NewKV(sr.client)
	_, errKP := kv.Delete(context.TODO(), key)
	return errKP
}

// 撤销租约
func (sr *ServiceReg) RevokeLease() error {

	log.Printf("[%d] start revoke service.", sr.leaseResp.ID)

	sr.canclefunc()
	time.Sleep(time.Duration(2) * time.Second)
	_, errSLR := sr.lease.Revoke(context.TODO(), sr.leaseResp.ID)
	return errSLR
}

func main() {
	ser, _ := NewServiceReg([]string{"10.122.105.131:12379", "10.122.105.131:22379", "10.122.105.131:32379"}, 6)
	ser.PutService("/node/111", "heiheihei")
	ser.PutService("/node/1111", "heiheihei2")
	ser.PutService("/node/11111", "heiheihei3")
	time.Sleep(time.Duration(3) * time.Second)

	ser.PutService("/node/1111", "heiheihei22")
	time.Sleep(time.Duration(3) * time.Second)

	ser.DeleteService("/node/11111")
	time.Sleep(time.Duration(3) * time.Second)

	ser2, _ := NewServiceReg([]string{"10.122.105.131:12379"}, 6)
	ser2.PutService("/node/222", "hahaha")
	time.Sleep(time.Duration(3) * time.Second)

	ser3, _ := NewServiceReg([]string{"10.122.105.131:12379"}, 6)
	ser3.PutService("/node/333", "wawawa")
	time.Sleep(time.Duration(5) * time.Second)

	//ser.RevokeLease()
	//ser2.RevokeLease()
	//ser3.RevokeLease()

	select {}

}
