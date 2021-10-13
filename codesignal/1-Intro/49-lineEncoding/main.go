package main

import "fmt"

func lineEncoding(s string) string {
	lastChar := rune(s[0])
	count := 0
	result := ""
	for _, char := range s + "\n" {
		if char == lastChar {
			count++
			lastChar = char
			continue
		}
		if count == 1 {
			result += string(lastChar)
		} else {
			result += fmt.Sprintf("%d%s", count, string(lastChar))
		}
		lastChar = char
		count = 1
	}
	return result
}

// Challenge:
// https://app.codesignal.com/arcade/intro/level-11/o2uq6eTuvk7atGadR
func main() {
	fmt.Println(lineEncoding("aabbbc")) //Should print "2a3bc"
}
