package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.EqualFold("v1.2.0", ""))
	fmt.Println(strings.EqualFold("v1.2.0", "V1.2.0"))
	fmt.Println(strings.EqualFold("v1.2.0.1", "V1.2.0"))
	fmt.Println(strings.EqualFold("v1.2.0.1", "v1.2.0"))
	fmt.Println(strings.EqualFold("V2.0.1.B04.0131", "V1.2.0"))
}
