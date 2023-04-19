package main

import "fmt"

func main() {
	switch 2 {
	case 1:
		fmt.Println(1)
	case 2:
		switch "3" {
		case "4":
			fmt.Println("4")
		case "5":
			fmt.Println("5")
		}
	default:
		fmt.Println(0)
	}
}
