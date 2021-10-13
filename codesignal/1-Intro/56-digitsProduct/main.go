package main

import (
	"fmt"
	"strconv"
)

func digitsProduct(product int) int {
	for i := 1; i <= 3558; i++ {
		p := 1
		for _, digit := range strconv.Itoa(i) {
			pInt, _ := strconv.Atoi(string(digit))
			p *= pInt
		}
		if p == product {
			return i
		}
	}
	return -1
}

// Challenge:
// https://app.codesignal.com/arcade/intro/level-12/NJJhENpgheFRQbPRA
func main() {
	fmt.Println(digitsProduct(12)) // Should print 26
	fmt.Println(digitsProduct(19)) // Should print -1
}
