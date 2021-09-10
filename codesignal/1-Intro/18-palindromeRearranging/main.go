package main

import "fmt"

func palindromeRearranging(inputString string) bool {
	counter := make(map[rune]int)
	for _, char := range inputString {
		counter[char]++
	}
	var oddKeys int
	for _, count := range counter {
		if count%2 == 0 {
			continue
		}
		oddKeys++
		if oddKeys > 1 {
			return false
		}
	}
	return true
}

// Challenge link
// https://app.codesignal.com/arcade/intro/level-4/Xfeo7r9SBSpo3Wico
func main() {
	fmt.Println(palindromeRearranging("aabb")) //Should print true
}
