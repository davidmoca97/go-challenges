package main

import "fmt"

func avoidObstacles(inputArray []int) int {
lengthLoop:
	for length := 2; length <= 1000; length++ {
		for _, obstacle := range inputArray {
			if obstacle%length == 0 {
				continue lengthLoop
			}
		}
		return length
	}
	return 1001
}

// Challenge:
// https://app.codesignal.com/arcade/intro/level-5/XC9Q2DhRRKQrfLhb5
func main() {
	input := []int{5, 3, 6, 7, 9}
	fmt.Println(avoidObstacles(input)) // Should print 4
}
