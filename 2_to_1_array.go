package main

import "fmt"

func main() {
	fmt.Println([][]string{{"a", "b", "c", "d", "e", "h"}, {"e", "f", "g", "h", "i"}, {"z"}})
	fmt.Println(stringArraysToArray([][]string{{"a", "b", "c", "d"}, {"e", "f", "g", "h", "i"}, {"z"}}))
}

func stringArraysToArray(arrays [][]string) []string {
	index := 0
	cnt := 0
	arrayCnt := len(arrays)
	for i := 0; i < arrayCnt; i++ {
		cnt += len(arrays[i])
	}
	outArray := make([]string, cnt)
	for i := 0; i < arrayCnt; i++ {
		cntI := len(arrays[i])
		for j := 0; j < cntI; j++ {
			outArray[index] = arrays[i][j]
			index++
		}
	}
	return outArray
}
