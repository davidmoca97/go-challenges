package main

import (
	"fmt"
)

func reverseInParentheses(inputString string) string {
	var stack []rune
	var result string
	for _, char := range inputString {
		if char == ')' {
			var subStr string
			for i := len(stack) - 1; i >= 0; i-- {
				charFromStack := stack[i]
				stack = stack[:i]
				if charFromStack == '(' {
					break
				}
				subStr += string(charFromStack)
			}
			if len(stack) > 0 {
				stack = append(stack, []rune(subStr)...)
				continue
			}
			result += subStr
			continue
		}
		if char == '(' || len(stack) > 0 {
			stack = append(stack, char)
			continue
		}
		result += string(char)
	}
	return result
}

func main() {
	res := reverseInParentheses("foo(bar(baz))blim")
	fmt.Println(res) // Should return foobazrabblim
}
