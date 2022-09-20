package main

import "fmt"

func main() {
	res := make([]byte, 0, 8)
	Per2(res)

	fmt.Println(res)
}

func Per2(res2 []byte) {
	res2 = append(res2, 'a')
}
