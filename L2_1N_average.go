package main

import (
	"fmt"
)

func main() {
	a := []byte{1, 2, 3, 4}
	b := [][]byte{{12, 7, 9, 1}}

	distance := getAverageFeatureL2Distance(a, b)
	fmt.Println(distance)

}

// 计算feature 的L2 距离算法
func getFeatureL2Distance(f1 []byte, f2 []byte) float64 {
	var dis float64
	for i, v := range f1 {
		square := float64((v - f2[i]) * (v - f2[i]))
		dis = dis + square
	}
	return dis
}

// 计算feature 与40张低质量图特征的平均数据
func getAverageFeatureL2Distance(f1 []byte, fs [][]byte) float64 {
	arr := make([]float64, 0, len(fs))
	for _, f := range fs {
		if len(f1) != len(f) { // feature 长度不同
			continue
		}
		dis := getFeatureL2Distance(f1, f)
		arr = append(arr, dis)
	}
	if len(arr) == 0 { // 非正常数据
		return -1
	}
	return getFloatArraySum(arr) / float64(len(arr))
}

// float 数组求和
func getFloatArraySum(array []float64) float64 {
	sum := 0.0
	for _, v := range array {
		sum += v
	}
	return sum
}
