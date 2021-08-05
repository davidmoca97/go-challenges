package main

import (
	"fmt"
	"regexp"
)

func firstDigit(inputString string) string {
	regex := regexp.MustCompile(`\d`)
	match := regex.Find([]byte(inputString))
	return string(match)
}

// Challenge
// https://app.codesignal.com/arcade/intro/level-8/rRGGbTtwZe2mA8Wov
func main() {
	fmt.Println(firstDigit("var_1__Int")) //Should print 1
}
