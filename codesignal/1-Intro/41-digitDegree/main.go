package main

import (
	"fmt"
	"strconv"
)

func digitDegree(n int) int {
	var times int
	for n > 9 {
		times++
		str := strconv.Itoa(n)
		var newNumber int
		for _, digitStr := range str {
			digit, _ := strconv.Atoi(string(digitStr))
			newNumber += digit
		}
		n = newNumber
	}
	return times
}

// Challenge:
// https://app.codesignal.com/arcade/intro/level-9/hoLtYWbjdrD2PF6yo/
func main() {
	fmt.Println(digitDegree(91)) //Should print 2
}
