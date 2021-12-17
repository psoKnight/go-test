package main

import "fmt"

func main() {
	var f32 float32 = 0.1111111111
	var f64 float32 = 0.2222222222

	fmt.Println(float32(f64))
	fmt.Println(float64(f32))
}
