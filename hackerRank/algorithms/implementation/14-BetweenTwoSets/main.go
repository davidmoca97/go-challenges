package main

func MCD(a, b int32) int32 {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func MCDSlice(s []int32) int32 {
	if len(s) == 0 {
		return 0
	}
	previous := s[0]
	for i := 1; i < len(s); i++ {
		previous = MCD(previous, s[i])
	}
	return previous
}

func MCM(a, b int32) int32 {
	return a * (b / MCD(a, b))
}

func MCMSlice(s []int32) int32 {
	if len(s) == 0 {
		return 0
	}
	previous := s[0]
	for i := 1; i < len(s); i++ {
		previous = MCM(previous, s[i])
	}
	return previous
}

/*
 * Complete the 'getTotalX' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

// Challenge:
// https://www.hackerrank.com/challenges/between-two-sets/problem
func getTotalX(a []int32, b []int32) int32 {
	// Write your code here
	mcm := MCMSlice(a)
	mcd := MCDSlice(b)

	var count int32
	for i := mcm; i <= mcd; i = i + mcm {
		if mcd%i == 0 {
			count++
		}
	}
	return count
}
