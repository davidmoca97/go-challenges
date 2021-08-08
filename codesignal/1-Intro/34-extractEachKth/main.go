package main

import "fmt"

func extractEachKth(inputArray []int, k int) []int {
	result := []int{}

	for i, item := range inputArray {
		if (i+1)%k > 0 {
			result = append(result, item)
		}
	}

	return result
}

// Challenge:
// https://app.codesignal.com/arcade/intro/level-8/3AgqcKrxbwFhd3Z3R
func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	k := 3
	fmt.Println(extractEachKth(slice, k)) // Should print [1, 2, 4, 5, 7, 8, 10]
}
