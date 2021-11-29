package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan bool)

	ticket := time.NewTimer(time.Second * 3)

	defer ticket.Stop()

	for {
		select {
		case <-ticket.C:
			fmt.Println("Time end.")
		case <-ch:
			fmt.Println("Channel output.")
		}

		break
	}

}
