package main

import "fmt"

func invertString(s string) string {
	inverse := ""
	for i := len(s) - 1; i >= 0; i-- {
		inverse += string(s[i])
	}
	return inverse
}

func buildPalindrome(st string) string {
	inverseString := invertString(st)
	missingSubstring := ""
	for idx := range st {
		currentResult := missingSubstring + inverseString
		if currentResult == invertString(currentResult) {
			return currentResult
		}
		missingSubstring += string(st[idx])
	}

	return missingSubstring + inverseString
}

// Challenge:
// https://app.codesignal.com/arcade/intro/level-10/ppZ9zSufpjyzAsSEx
func main() {
	fmt.Println(buildPalindrome("abcdc")) //Should print abcdcba
}
