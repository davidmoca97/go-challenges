package main

import (
	"fmt"
)

func digitDegree(n int) int {
	var times int
	for n > 9 {
		times++
		copyOfN := n
		var newNumber = 0
		for copyOfN > 0 {
			digit := copyOfN % 10
			newNumber += digit
			copyOfN = copyOfN / 10
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
