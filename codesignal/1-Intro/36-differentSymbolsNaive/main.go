package main

import "fmt"

func differentSymbolsNaive(s string) int {
	dictionary := make(map[rune]bool)
	for _, char := range s {
		dictionary[char] = true
	}
	return len(dictionary)
}

// Challenge:
// https://app.codesignal.com/arcade/intro/level-8/8N7p3MqzGQg5vFJfZ
func main() {
	fmt.Println(differentSymbolsNaive("cabca")) //should print 3
}
