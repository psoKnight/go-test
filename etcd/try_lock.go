package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
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
		Endpoints:   []string{"10.122.105.131:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	lockKey := "/lock"

	go func() {

		log.Println("start go3")

		session, err := concurrency.NewSession(cli)
		if err != nil {
			log.Fatal(err)
		}
		m := concurrency.NewMutex(session, lockKey)
		if err := m.TryLock(context.TODO()); err != nil {
			log.Fatal("go3 try get mutex failed " + err.Error())
		}
		fmt.Printf("go3 try get mutex success\n")
		fmt.Println(m)
		//time.Sleep(time.Duration(15) * time.Second)
		m.Unlock(context.TODO())
		fmt.Printf("go3 release lock\n")
	}()

	<-c

}
