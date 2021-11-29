package main

import (
	"fmt"
	"strings"
)

func main() {
	videoId := strings.Split("1_2#3[456]478", "[")[0] // 02 core format is {videoId}[sth]
	videoId = strings.Split(videoId, "#")[0]          // 05 core format is {videoId}#sth
	videoId = strings.Split(videoId, "_")[0]          // 08 core format is {videoId}_sth

	fmt.Println(videoId)
}
