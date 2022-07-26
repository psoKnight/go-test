package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"math"
)

func main() {
	arrA := []float64{1, 2, 3, 4}
	arrB := []float64{12, 7, 9, 1}

	if len(arrA) != len(arrB) {
		logrus.Errorf("array length is not match")
		return
	}

	var sum float64
	for i, v := range arrA {
		square := math.Pow(v-arrB[i], 2)
		sum = sum + square
	}

	fmt.Println(sum)
}
