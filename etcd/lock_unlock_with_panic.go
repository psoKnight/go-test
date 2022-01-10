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
		Endpoints:   []string{"10.122.105.131:12379"},
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

		session, err := concurrency.NewSession(cli, concurrency.WithTTL(10))
		if err != nil {
			log.Fatal(err)
		}
		m := concurrency.NewMutex(session, lockKey)
		if err := m.Lock(context.TODO()); err != nil {
			log.Fatal("go1 get mutex failed " + err.Error())
		}
		fmt.Printf("go1 get mutex success\n")
		fmt.Println(m)

		time.Sleep(time.Duration(15) * time.Second)
		fmt.Println(time.Now().Format("2006-01-02 15:04:05.000000000"))

		panic("oh,it's panic")

		m.Unlock(context.TODO())
		fmt.Printf("go1 release lock\n")
	}()

	<-c
}
