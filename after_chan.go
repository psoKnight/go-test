package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	a := time.After(3 * time.Second)
	fmt.Println(<-a)
	fmt.Println(a)
}
