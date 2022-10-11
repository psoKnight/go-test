package main

import "fmt"

func main() {
	lists := make([]int, 0)
	lists = append(lists, 1)

	for _, num := range lists[1:] {
		fmt.Println(num)
	}
}
