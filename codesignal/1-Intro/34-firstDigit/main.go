package main

import (
	"fmt"
	"regexp"
)

func firstDigit(inputString string) string {
	return regexp.MustCompile(`\d`).FindString(inputString)
}

// Challenge
// https://app.codesignal.com/arcade/intro/level-8/rRGGbTtwZe2mA8Wov
func main() {
	fmt.Println(firstDigit("var_1__Int")) //Should print 1
}
