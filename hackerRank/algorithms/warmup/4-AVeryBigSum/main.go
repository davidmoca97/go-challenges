package main

import "fmt"

// Complete the aVeryBigSum function below.
func aVeryBigSum(ar []int64) int64 {
	var total int64
	for _, number := range ar {
		total += number
	}
	return total
}

// Challenge:
// https://www.hackerrank.com/challenges/a-very-big-sum
func main() {
	numbers := []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	fmt.Println(aVeryBigSum(numbers)) // Should print 5000000015
}
