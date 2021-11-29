package main

import (
	"fmt"
)

func main() {
	type BatchPersonRes struct {
		remark *string
	}

	fmt.Println("before:", BatchPersonRes{})

	var r = "rem"

	batchPersonRes := BatchPersonRes{remark: &r}

	fmt.Println("after:", batchPersonRes)

}
