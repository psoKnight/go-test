package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func main() {
	str := "http://169.42.42.42:8333/megconnect_face_full/702BBE60-2EB3-DF51-B80A-B15C3F787FDD.jpg"

	if str == "" {
		return
	}

	splits := strings.Split(str, "/")
	if len(splits) != 5 {
		return
	}

	encode := base64.StdEncoding.EncodeToString([]byte("cactus100_" + splits[3]))

	uri := "_" + encode + "_" + splits[4]

	fmt.Println(uri)
}
