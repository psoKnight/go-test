package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 500)

	for t := range ticker.C {
		fmt.Println("Tick at", t)
	}

}
