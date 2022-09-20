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
		Endpoints:   []string{"10.117.49.69:22379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	lockKey := "/lock"

	go func() {

		log.Println("start go4")

		log.Println(cli.MemberList(context.TODO()))

		session, err := concurrency.NewSession(cli)
		if err != nil {
			log.Fatal(err)
		}
		m := concurrency.NewMutex(session, lockKey)
		if err := m.Lock(context.TODO()); err != nil {
			log.Fatal("go4 get mutex failed " + err.Error())
		}

		fmt.Println(time.Now().Format("2006-01-02 15:04:05.000000000"))

		//time.Sleep(time.Duration(10) * time.Second)

		fmt.Printf("go4 get mutex success\n")
		fmt.Println(m)

		m.Unlock(context.TODO())
		fmt.Printf("go4 release lock\n")
	}()

	go func() {

		log.Println("start go5")

		session, err := concurrency.NewSession(cli)
		if err != nil {
			log.Fatal(err)
		}
		m := concurrency.NewMutex(session, lockKey)
		if err := m.Lock(context.TODO()); err != nil {
			log.Fatal("go5 get mutex failed " + err.Error())
		}

		fmt.Println(time.Now().Format("2006-01-02 15:04:05.000000000"))

		fmt.Printf("go5 get mutex success\n")
		fmt.Println(m)
		m.Unlock(context.TODO())
		fmt.Printf("go5 release lock\n")
	}()

	<-c
}
