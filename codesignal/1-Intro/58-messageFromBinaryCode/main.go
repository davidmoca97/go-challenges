package main

import (
	"fmt"
	"strconv"
)

func messageFromBinaryCode(code string) string {
	rightLimit := 8
	var result []byte
	for rightLimit < len(code)+1 {
		bytee := code[rightLimit-8 : rightLimit]
		b, _ := strconv.ParseInt(bytee, 2, 64)
		result = append(result, byte(b))
		rightLimit += 8
	}

	return string(result)
}

// Challenge:
// https://app.codesignal.com/arcade/intro/level-12/sCpwzJCyBy2tDSxKW
func main() {
	msg := "010010000110010101101100011011000110111100100001"
	fmt.Println(messageFromBinaryCode(msg)) //Should print Hello!
}
