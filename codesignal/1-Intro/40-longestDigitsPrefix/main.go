package main

import (
	"fmt"
	"regexp"
)

func longestDigitsPrefix(inputString string) string {
	matches := regexp.MustCompile(`^\d+`).FindAllString(inputString, -1)
	if len(matches) == 0 {
		return ""
	}
	return matches[0]
}

// Challenge:
// https://app.codesignal.com/arcade/intro/level-9/AACpNbZANCkhHWNs3
func main() {
	fmt.Println(longestDigitsPrefix("123aa1")) //Should print 123
}
