package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	contains := strings.Contains(errors.New("this is an er/ror").Error(), "error")

	if contains {
		fmt.Println(contains)
	}
}
