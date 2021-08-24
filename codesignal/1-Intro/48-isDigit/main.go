package main

import (
	"fmt"
	"regexp"
)

func isDigit(symbol string) bool {
	return regexp.MustCompile(`^\d$`).MatchString(symbol)
}

// Challenge:
// https://app.codesignal.com/arcade/intro/level-11/Zr2XXEpkBPtYWqPJu
func main() {
	fmt.Println(isDigit("0")) // Should print true
	fmt.Println(isDigit("-")) // Should print false
}
