package main

import "fmt"

func isBeautifulString(inputString string) bool {
	countMap := make(map[rune]int)
	for _, char := range inputString {
		countMap[char]++
	}

	abecedary := "abcdefghijklmnopqrstuvwxyz"
	for idx := range abecedary {
		if idx == 0 {
			continue
		}
		letter := rune(abecedary[idx])
		previousLetter := rune(abecedary[idx-1])
		if countMap[letter] > countMap[previousLetter] {
			return false
		}
	}
	return true
}

// Challenge:
// https://app.codesignal.com/arcade/intro/level-10/PHSQhLEw3K2CmhhXE
func main() {
	fmt.Println(isBeautifulString("bbbaacdafe")) //Should print true
}
