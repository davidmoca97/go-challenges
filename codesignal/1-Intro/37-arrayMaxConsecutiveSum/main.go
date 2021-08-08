package main

import (
	"fmt"
	"math"
)

func sum(slice []int) int {
	var total int
	for _, n := range slice {
		total += n
	}
	return total
}

func arrayMaxConsecutiveSum(inputArray []int, k int) int {
	max := math.MinInt32
	for i := 0; i <= len(inputArray)-k; i++ {
		sum := sum(inputArray[i : i+k])
		if sum > max {
			max = sum
		}
	}
	return max
}

// Challenge:
// https://app.codesignal.com/arcade/intro/level-8/Rqvw3daffNE7sT7d5
func main() {
	input := []int{1, 3, 2, 4}
	fmt.Println(arrayMaxConsecutiveSum(input, 3)) //Should print 9
}
