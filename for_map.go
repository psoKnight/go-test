package main

import "fmt"

func main() {
	var m map[string]interface{}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
