package main

import (
	"fmt"
	"regexp"
)

func longestDigitsPrefix(inputString string) string {
	matches := regexp.MustCompile(`^\d+`).FindAllString(inputString, -1)
	longest := ""
	for _, match := range matches {
		if len(match) > len(longest) {
			longest = match
		}
	}
	return longest
}

// Challenge:
// https://app.codesignal.com/arcade/intro/level-9/AACpNbZANCkhHWNs3
func main() {
	fmt.Println(longestDigitsPrefix("123aa1")) //Should print 123
}
