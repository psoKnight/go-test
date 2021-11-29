package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)

	num := 0

	go func() {
		for {
			<-ch
		}
	}()

	time.Sleep(time.Duration(1) * time.Second)

	f := func() {
		select {
		case ch <- true:
			fmt.Println("write")
		default:
			fmt.Println("default")
		}
	}

	for {
		go f()

		num++

		time.Sleep(time.Duration(500) * time.Millisecond)

		if num == 3 {
			break
		}
	}

}
