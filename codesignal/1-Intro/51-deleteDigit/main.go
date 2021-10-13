package main

import (
	"fmt"
	"strconv"
)

func deleteDigit(n int) int {
	str := strconv.Itoa(n)
	max := 0
	for idx := range str {
		newNumber, _ := strconv.Atoi(str[0:idx] + str[idx+1:])
		if newNumber > max {
			max = newNumber
		}
	}
	return max
}

// Challenge:
// https://app.codesignal.com/arcade/intro/level-11/vpfeqDwGZSzYNm2uX
func main() {
	fmt.Println(deleteDigit(152))  // Should print 52
	fmt.Println(deleteDigit(1001)) // Should print 101
}
