package main

import (
	"math/rand"
	"os"
	"runtime/pprof"
	"time"
)

func generate(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}

/**
方式一：$ go tool pprof -http=:9999 cpu.pprof

方式二：
$ go tool pprof cpu.pprof
Type: cpu
Time: Nov 11, 2022 at 1:48pm (CST)
Duration: 13.93s, Total samples = 7.62s (54.71%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 7.37s, 96.72% of 7.62s total
Dropped 30 nodes (cum <= 0.04s)
Showing top 10 nodes out of 21
      flat  flat%   sum%        cum   cum%
     7.23s 94.88% 94.88%      7.25s 95.14%  main.bubbleSort (inline)
     0.04s  0.52% 95.41%      0.04s  0.52%  runtime.stdcall1
     0.03s  0.39% 95.80%      0.15s  1.97%  runtime.findrunnable
     0.02s  0.26% 96.06%      0.22s  2.89%  runtime.schedule
     0.02s  0.26% 96.33%      0.06s  0.79%  runtime/pprof.(*profileBuilder).addCPUData
     0.01s  0.13% 96.46%      0.04s  0.52%  runtime.checkTimers
     0.01s  0.13% 96.59%      0.08s  1.05%  runtime.stealWork
     0.01s  0.13% 96.72%      0.04s  0.52%  runtime/pprof.(*profMap).lookup
         0     0% 96.72%      7.25s 95.14%  main.main
         0     0% 96.72%      7.25s 95.14%  runtime.main
(pprof) top --cum
Showing nodes accounting for 7.31s, 95.93% of 7.62s total
Dropped 30 nodes (cum <= 0.04s)
Showing top 10 nodes out of 21
      flat  flat%   sum%        cum   cum%
     7.23s 94.88% 94.88%      7.25s 95.14%  main.bubbleSort (inline)
         0     0% 94.88%      7.25s 95.14%  main.main
         0     0% 94.88%      7.25s 95.14%  runtime.main
         0     0% 94.88%      0.26s  3.41%  runtime.mcall
         0     0% 94.88%      0.26s  3.41%  runtime.park_m
     0.02s  0.26% 95.14%      0.22s  2.89%  runtime.schedule
     0.03s  0.39% 95.54%      0.15s  1.97%  runtime.findrunnable
         0     0% 95.54%      0.11s  1.44%  runtime/pprof.profileWriter
     0.01s  0.13% 95.67%      0.08s  1.05%  runtime.stealWork
     0.02s  0.26% 95.93%      0.06s  0.79%  runtime/pprof.(*profileBuilder).addCPUData
(pprof)

*/
func main() {
	f, _ := os.OpenFile("./pprof/cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	n := 10
	for i := 0; i < 5; i++ {
		nums := generate(n)
		bubbleSort(nums)
		n *= 10
	}
}
