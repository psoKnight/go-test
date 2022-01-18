package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c)

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"10.122.105.131:12379", "10.122.105.131:22379", "10.122.105.131:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	lockKey := "/lock"

	go func() {

		log.Println("start go1")

		log.Println(cli.MemberList(context.TODO()))

		session, err := concurrency.NewSession(cli, concurrency.WithTTL(1))
		if err != nil {
			log.Fatal(err)
		}
		m := concurrency.NewMutex(session, lockKey)
		if err := m.Lock(context.TODO()); err != nil {
			log.Fatal("go1 get mutex failed " + err.Error())
		}
		fmt.Printf("go1 get mutex success\n")
		/**

		my_mutex 工作原理
			每个mutex 会创建一个租约lease，并且租约是长期有效的
			使用prefix/leaseId 作为key 向etcd 插入数据
			etcd 确保每次插入的key 都有一个位移的CreateRevision，并且CreateRevision 是从 0 开始递增的
			第 2 条规定所有mutex 按照prefix/leaseId 作为key 向etcd 插入数据
			每个mutex 监视prefix 下所有key,并且获得这些key 的CreateRevision
			如果mutex 发现自己的prefix/leaseId 的CreateRevision 是这些prefix 下所有key 中最小的，那么当前mutex 获得锁

		type Session struct {
		    client *v3.Client
		    opts   *sessionOptions
		    id     v3.LeaseID

		    cancel context.CancelFunc
		    donec  <-chan struct{}
		}

		type Mutex struct {
		    s *Session  // 上面的Session struct

		    pfx   string  // 前缀
		    myKey string  // key，格式：prefix/leaseId
		    myRev int64   // Revision
		    hdr   *pb.ResponseHeader
		}


		*/
		fmt.Println(m)

		time.Sleep(time.Duration(15) * time.Second)
		fmt.Println(time.Now().Format("2006-01-02 15:04:05.000000000"))

		panic("oh,it's panic")

		m.Unlock(context.TODO())
		fmt.Printf("go1 release lock\n")
	}()

	<-c
}
