package main

import "fmt"

func main() {
	var f32 float32 = 0.1111
	var f64 float32 = 0.22222

	fmt.Println(float32(f64))
	fmt.Println(float64(f32))
}
