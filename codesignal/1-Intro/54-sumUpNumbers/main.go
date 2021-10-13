package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func sumUpNumbers(inputString string) int {
	matches := regexp.MustCompile(`\d+`).FindAllString(inputString, -1)
	sum := 0
	for _, numberStr := range matches {
		number, _ := strconv.Atoi(numberStr)
		sum += number
	}
	return sum
}

// Challenge:
// https://app.codesignal.com/arcade/intro/level-12/YqZwMJguZBY7Hz84T
func main() {
	fmt.Println(sumUpNumbers("2 apples, 12 oranges"))           //Should print 14
	fmt.Println(sumUpNumbers("Your payment method is invalid")) //Should print 0
}
