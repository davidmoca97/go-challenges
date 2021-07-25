package main

import (
	"fmt"
	"math"
)

func depositProfit(deposit int, rate int, threshold int) int {
	// y = d(r^x) -> x = logr(y/d)
	logX := func(x float64, v float64) float64 {
		return math.Log(v) / math.Log(x)
	}
	log := logX(float64(rate)/100+1, float64(threshold)/float64(deposit))
	return int(math.Ceil(log))
}

// Challenge
// https://app.codesignal.com/arcade/intro/level-7/8PxjMSncp9ApA4DAb
func main() {
	fmt.Println(depositProfit(100, 20, 170)) //Should print 3
}
