package main

import "fmt"

func main() {
	s := make([]int, 10)

	for i := 0; i < 10; i++ {
		s[i] = i + 1
	}

	fmt.Println(s)

	fS := s[0:5]
	bS := s[5:]

	fmt.Println("Before append:", fS)
	fmt.Println("Before append:", bS)

	fS = append(fS, 99)

	fmt.Println("After append:", fS)
	fmt.Println("After append:", bS)

}
