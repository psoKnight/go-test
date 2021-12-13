package main

import "fmt"

func main() {
	t := 1

	switch t {
	case 1:
		break
		fmt.Println(1) // break
	case 2:
		fmt.Println(2)
	default:
		fmt.Println(3)
	}
}
