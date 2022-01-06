package main

import (
	"fmt"
	"time"
)

func main() {

	tm := time.Unix(time.Now().Unix(), 0).Format("20060102030405")

	fmt.Println(tm)

}
