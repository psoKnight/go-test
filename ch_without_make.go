package main

import "fmt"

func main() {
	var ch chan int

	go writeCh1(ch)

	fmt.Println(<-ch)
}

func writeCh1(ch chan int) {
	ch <- 9
}
