package main

import (
	"fmt"
	"sort"
)

func miniMaxSum(arr []int32) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	var max, min int64
	for i := 0; i < 4; i++ {
		min += int64(arr[i])
		max += int64(arr[len(arr)-1-i])
	}
	fmt.Printf("%d %d\n", min, max)
}

// Challenge:
// https://www.hackerrank.com/challenges/mini-max-sum/problem
func main() {
	// Should print "16 24"
	miniMaxSum([]int32{1, 3, 5, 7, 9})
}
