package main

import "fmt"

func alphabeticShift(inputString string) string {
	var output string
	for _, char := range inputString {
		if char == 'z' {
			output += "a"
			continue
		}
		output += string(char + 1)
	}
	return output
}

// Challenge:
// https://app.codesignal.com/arcade/intro/level-6/PWLT8GBrv9xXy4Dui
func main() {
	str := "crazy"
	fmt.Println(alphabeticShift(str)) //Should print dsbaz
}
