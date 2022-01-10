package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type ConfigAtomic struct {
	a []int
}

func main() {
	var v atomic.Value

	// 写入数据
	go func() {
		var i int
		for {
			i++
			cfg := &ConfigAtomic{
				a: []int{i, i + 1, i + 2, i + 3, i + 4, i + 5},
			}
			v.Store(cfg)
		}
	}()

	// 读取数据
	var wg sync.WaitGroup
	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func() {
			for n := 0; n < 100; n++ {
				cfg := v.Load()
				fmt.Printf("%#v\n", cfg)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
