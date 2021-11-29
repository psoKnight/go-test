package main

import (
	"fmt"
	"time"
)

func main() {
	ticket := time.NewTicker(time.Second * time.Duration(1))
	defer ticket.Stop()

	num := 0

	for range ticket.C {

		fmt.Println(num)

		num++

	}

}
