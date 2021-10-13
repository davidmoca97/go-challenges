package main

import "fmt"

/*
 * Complete the simpleArraySum function below.
 */
func simpleArraySum(ar []int32) int32 {
	var total int32
	for _, num := range ar {
		total += num
	}
	return total
}

// Challenge:
// https://www.hackerrank.com/challenges/simple-array-sum/problem
func main() {
	fmt.Println(simpleArraySum([]int32{1, 2, 3, 4, 10, 11})) // Should print 31
}
