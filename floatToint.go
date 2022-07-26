package main

import (
	"fmt"
	"time"
)

func main() {
	var f float64 = 3.1415926

	fmt.Println(int64(f))
	fmt.Println(int64(time.Duration(2160).Seconds()))
}
