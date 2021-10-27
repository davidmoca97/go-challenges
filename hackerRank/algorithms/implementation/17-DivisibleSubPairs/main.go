package main

import (
	"fmt"
)

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	var count int32
	for i := 0; i < len(ar)-1; i++ {
		for j := i + 1; j < len(ar); j++ {
			if (ar[i]+ar[j])%k == 0 {
				count++
			}
		}
	}
	return count
}

// Challenge:
// https://www.hackerrank.com/challenges/divisible-sum-pairs/problem
func main() {
	// Should print 3
	fmt.Println(divisibleSumPairs(6, 5, []int32{1, 2, 3, 4, 5, 6}))
}
