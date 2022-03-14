package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mu := make(map[string]int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			time.Sleep(1 * time.Second)
			if _, ok := mu["g1"]; ok {
				fmt.Print(1)
				delete(mu, "g1")
				mu["g2"] = 1
			}
		}
	}()
	go func() {
		for {
			time.Sleep(1 * time.Second)
			if _, ok := mu["g2"]; ok {
				fmt.Print(2)
				delete(mu, "g2")
				mu["g3"] = 1
			}
		}
	}()
	go func() {
		for {
			time.Sleep(1 * time.Second)
			if _, ok := mu["g3"]; ok {
				fmt.Print(3)
				delete(mu, "g3")
				mu["g1"] = 1
			}
		}
	}()

	mu["g1"] = 1
	wg.Wait()
}
