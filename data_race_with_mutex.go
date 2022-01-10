package main

import (
	"fmt"
	"sync"
)

type ConfigMutex struct {
	a []int
}

func main() {
	cfg := &ConfigMutex{}
	var mux sync.RWMutex

	// 启动一个 writer goroutine，不断写入数据
	go func() {
		i := 0

		for {
			i++
			// 进行数据写入时，先通过锁进行锁定
			mux.Lock()
			cfg.a = []int{i, i + 1, i + 2, i + 3, i + 4, i + 5}
			mux.Unlock()
		}
	}()

	// 启动多个 reader goroutine，不断获取数据
	var wg sync.WaitGroup
	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func() {
			for n := 0; n < 100; n++ {
				// 因为这里只是需要读取数据，所以只需要加一个读锁即可
				mux.RLock()
				fmt.Printf("%#v\n", cfg)
				mux.RUnlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
