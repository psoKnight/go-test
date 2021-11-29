package main

import (
	"fmt"
	"time"
)

func main() {

	// 协程A
	go func() {
		for {
			fmt.Println("goroutine1_print")
		}
	}()

	// 协程B
	go func() {
		time.Sleep(1 * time.Second)
		panic("goroutine2_panic")
	}()

	time.Sleep(2 * time.Second)
}
