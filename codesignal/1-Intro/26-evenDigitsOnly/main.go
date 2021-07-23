package main

import (
	"fmt"
)

func evenDigitsOnly(n int) bool {
	divider := 10
	for n > 0 {
		digit := n % divider
		if digit%2 != 0 {
			return false
		}
		n /= divider
	}
	return true
}

// Challenge link:
// https://app.codesignal.com/arcade/intro/level-6/6cmcmszJQr6GQzRwW
func main() {
	fmt.Println(evenDigitsOnly(248622)) //Should return true
}
