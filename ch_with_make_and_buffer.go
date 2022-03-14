package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	go writeCh3(ch)

	fmt.Println(<-ch)
}

func writeCh3(ch chan int) {
	ch <- 9
}
