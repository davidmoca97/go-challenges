package main

import "fmt"

func circleOfNumbers(n int, firstNumber int) int {
	return (n/2 + firstNumber) % n
}

// Challenge:
// https://app.codesignal.com/arcade/intro/level-7/vExYvcGnFsEYSt8nQ
func main() {
	fmt.Println(circleOfNumbers(10, 2)) //Should print 7
}
