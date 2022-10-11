package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Format(time.RFC3339), time.Now().Unix())
}
