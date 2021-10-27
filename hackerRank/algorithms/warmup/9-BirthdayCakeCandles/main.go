package main

import (
	"fmt"
	"math"
)

func birthdayCakeCandles(candles []int32) int32 {
	max := int32(math.MinInt32)
	var maxCount int32
	for _, candle := range candles {
		if max == candle {
			maxCount++
			continue
		}
		if candle > max {
			max = candle
			maxCount = 1
		}
	}
	return maxCount
}

// Challenge:
// https://www.hackerrank.com/challenges/birthday-cake-candles/problem
func main() {
	fmt.Println(birthdayCakeCandles([]int32{4, 4, 1, 3}))
}
