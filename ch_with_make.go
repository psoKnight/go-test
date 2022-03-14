package main

import "fmt"

func main() {
	ch := make(chan int)

	go writeCh2(ch)

	fmt.Println(<-ch)
}

func writeCh2(ch chan int) {
	ch <- 9
}
