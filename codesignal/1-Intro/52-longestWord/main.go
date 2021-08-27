package main

import (
	"fmt"
	"regexp"
)

func longestWord(text string) string {
	matches := regexp.MustCompile(`(?i)[a-z]+`).FindAllString(text, -1)
	idxOfMax := 0

	for idx := range matches {
		if len(matches[idx]) > len(matches[idxOfMax]) {
			idxOfMax = idx
		}
	}
	return matches[idxOfMax]
}

// Challenge:
// https://app.codesignal.com/arcade/intro/level-12/s9qvXv4yTaWg8g4ma
func main() {
	fmt.Println(longestWord("Ready, steady, go!")) //Should print steady
}
