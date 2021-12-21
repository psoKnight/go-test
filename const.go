package main

import "fmt"

func main() {

	const (
		A int32 = iota
		B
		C
		D
		E
		F       = 99
		G       = 100
		H int32 = iota + 2
		I
		J
	)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Println(G)
	fmt.Println(H)
	fmt.Println(I)
	fmt.Println(J)

	fmt.Println("")

	const (
		HatColorUnknow int = iota + 16
		HatColorWhite
		HatColorOther
	)

	fmt.Println(HatColorUnknow, HatColorWhite, HatColorOther)

}
