package main

import (
	"fmt"
	"time"
)

func main() {
	var ttl int64 = 5
	fmt.Println((time.Duration(ttl) * time.Second) / 3.0)
}
