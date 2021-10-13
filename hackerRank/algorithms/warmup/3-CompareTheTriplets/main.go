package main

import "fmt"

// Complete the compareTriplets function below.
func compareTriplets(a []int32, b []int32) []int32 {
	returnPoints := []int32{0, 0}
	for i := 0; i < len(b); i++ {
		if a[i] > b[i] {
			returnPoints[0]++
		} else if a[i] < b[i] {
			returnPoints[1]++
		}
	}
	return returnPoints
}

// Challenge:
// https://www.hackerrank.com/challenges/compare-the-triplets/problem
func main() {
	a := []int32{5, 6, 7}
	b := []int32{3, 6, 10}
	fmt.Println(compareTriplets(a, b)) //Should print [1, 1]
}
