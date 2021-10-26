package main

import "fmt"

func plusMinus(arr []int32) {
	var positives, negatives, zeros int64
	var arrLen = float64(len(arr))
	for _, v := range arr {
		if v == 0 {
			zeros++
			continue
		}
		if v > 0 {
			positives++
			continue
		}
		negatives++
	}
	fmt.Printf("%.6f\n", float64(positives)/arrLen)
	fmt.Printf("%.6f\n", float64(negatives)/arrLen)
	fmt.Printf("%.6f\n", float64(zeros)/arrLen)
}

// Challenge:
// https://www.hackerrank.com/challenges/plus-minus/problem
func main() {
	/*
		Should print:
		0.400000
		0.400000
		0.200000
	*/
	plusMinus([]int32{1, 1, 0, -1, -1})
}
