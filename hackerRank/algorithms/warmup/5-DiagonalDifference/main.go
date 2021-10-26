package main

import (
	"fmt"
	"math"
)

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func diagonalDifference(arr [][]int32) int32 {
	var leftToRight, rightToLeft int32
	sliceLength := len(arr) - 1

	for i := 0; i < len(arr); i++ {
		leftToRight += arr[i][i]
		rightToLeft += arr[i][sliceLength-i]
	}

	return int32(math.Abs(float64(leftToRight - rightToLeft)))
}

// Challenge:
// https://www.hackerrank.com/challenges/diagonal-difference/problem
func main() {
	res := diagonalDifference([][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	})
	fmt.Println(res) // Should print 2
}
