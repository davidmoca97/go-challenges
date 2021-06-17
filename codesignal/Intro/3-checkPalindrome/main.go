package main

import "fmt"

func checkPalindromeWord(inputString string) bool {
	if inputString == "" {
		return false
	}

	reverseString := ""
	for _, char := range inputString {
		reverseString = string(char) + reverseString
	}
	return reverseString == inputString
}

func main() {
	fmt.Println(checkPalindromeWord("anitalavalatina"))
	fmt.Println(checkPalindromeWord("lol"))
	fmt.Println(checkPalindromeWord("nosoypalindroma"))
}
