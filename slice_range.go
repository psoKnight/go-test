package main

import (
	"fmt"
)

func main() {
	// 示例1
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	all := make([]*int, 0)
	for _, item := range items {
		all = append(all, &item)
	}
	for _, a := range all {
		fmt.Println(fmt.Sprintf("%d\n", *a))
	}

	// 示例2
	items2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	all2 := make([]*int, 0)
	for _, item2 := range items2 {
		it := item2
		all2 = append(all2, &it)
	}
	for _, a2 := range all2 {
		fmt.Println(fmt.Sprintf("%d\n", *a2))
	}

	// 示例3
	items3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	all3 := make([]int, 0)
	for _, item3 := range items3 {
		all3 = append(all3, item3)
	}
	for _, a3 := range all3 {
		fmt.Println(fmt.Sprintf("%d\n", a3))
	}

	// 示例4
	items4 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	all4 := make([]int, 0)
	for _, item4 := range items4 {
		all4 = append(all4, item4)
	}
	for _, a4 := range all4 {
		go func() {
			fmt.Println(fmt.Sprintf("%d\n", a4))
		}()
	}
}
