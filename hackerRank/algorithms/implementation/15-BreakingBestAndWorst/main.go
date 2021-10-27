package main

import "fmt"

/*
 * Complete the 'breakingRecords' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY scores as parameter.
 */

func breakingRecords(scores []int32) []int32 {
	max, min := scores[0], scores[0]
	var brokeMaxCount, brokeMinCount int32
	for i := 1; i < len(scores); i++ {
		if scores[i] > max {
			max = scores[i]
			brokeMaxCount++
			continue
		}
		if scores[i] < min {
			min = scores[i]
			brokeMinCount++
		}
	}
	return []int32{brokeMaxCount, brokeMinCount}
}

// Challenge:
// https://www.hackerrank.com/challenges/breaking-best-and-worst-records/problem
func main() {
	// Should print [2, 4]
	fmt.Println(breakingRecords([]int32{10, 5, 20, 20, 4, 5, 2, 25, 1}))
}
