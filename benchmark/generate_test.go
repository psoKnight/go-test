package benchmark

import (
	"math/rand"
	"testing"
	"time"
)

/**
go test -bench .
go test -benchmem -bench .

基准测试报告每一列值对应的含义如下：
type BenchmarkResult struct {
    N         int           // 迭代次数
    T         time.Duration // 基准测试花费的时间
    Bytes     int64         // 一次迭代处理的字节数
    MemAllocs uint64        // 总的分配内存的次数
    MemBytes  uint64        // 总的分配内存的字节数
}

goos: windows
goarch: amd64
pkg: go-test/benchmark
cpu: Intel(R) Core(TM) i5-8265U CPU @ 1.60GHz
BenchmarkGenerate1000-8            37124             34983 ns/op
BenchmarkGenerate10000-8            4417            269642 ns/op
BenchmarkGenerate100000-8            378           2814750 ns/op
BenchmarkGenerate1000000-8            38          26822171 ns/op
PASS
ok      go-test/benchmark       9.989s
*/

func generate(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}
func benchmarkGenerate(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		generate(i)
	}
}

func BenchmarkGenerate1000(b *testing.B)    { benchmarkGenerate(1000, b) }
func BenchmarkGenerate10000(b *testing.B)   { benchmarkGenerate(10000, b) }
func BenchmarkGenerate100000(b *testing.B)  { benchmarkGenerate(100000, b) }
func BenchmarkGenerate1000000(b *testing.B) { benchmarkGenerate(1000000, b) }
