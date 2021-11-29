package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("v1.2.0" > "v1.3.0") // false

	major, minor, build, err := GetVersionSplit("v1.2.0")
	if err != nil {
		return
	}

	if major >= 1 || (minor >= 1 && build >= 2) {
		fmt.Println("Above true.")
	} else {
		fmt.Println("Above false.")
	}

}

func GetVersionSplit(s string) (int, int, int, error) {

	if s == "" {
		return 0, 0, 0, errors.New("String is empty.")
	}

	sArr := strings.Split(s, ".")
	if len(sArr) < 3 {
		return 0, 0, 0, errors.New("The string length is incorrect.")
	}

	major, err := strconv.Atoi(sArr[0][1:])
	if err != nil {
		return 0, 0, 0, err
	}
	minor, err := strconv.Atoi(sArr[1])
	if err != nil {
		return 0, 0, 0, err
	}
	build, err := strconv.Atoi(sArr[2])
	if err != nil {
		return 0, 0, 0, err
	}

	return major, minor, build, nil
}
