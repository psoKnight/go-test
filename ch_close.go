package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		time.Sleep(time.Hour)
	}()
	c := make(chan int, 10)
	c <- 1
	c <- 2
	close(c)
	//c <- 3

	//fmt.Println(<-c) //1
	//fmt.Println(<-c) //2
	//fmt.Println(<-c) //0
	//fmt.Println(<-c) //0

	for i := range c {
		fmt.Println(i)
	}
}
