package main

import (
	"fmt"
	"sort"
)

func factorial(n int) int64 {
	res := int64(n)
	for i := n - 1; i >= 2; i-- {
		res *= int64(i)
	}
	return res
}

// Used Heaps's non-recursive algorithm to get the permutations
// https://en.wikipedia.org/wiki/Heap%27s_algorithm
func newPermutationsGenerator(arr []string) func() ([]string, bool) {
	c := make([]int, len(arr))
	arrCopy := make([]string, len(arr))
	copy(arrCopy, arr)

	calls := 0
	i := 0
	totalPermutations := factorial(len(arr))
	return func() ([]string, bool) {
		calls++
		hasMorePermutations := int64(calls) < totalPermutations
		if calls == 1 {
			return arrCopy, hasMorePermutations
		}
		for i < len(arrCopy) {
			if c[i] < i {
				if i%2 == 0 {
					arrCopy[0], arrCopy[i] = arrCopy[i], arrCopy[0]
				} else {
					arrCopy[c[i]], arrCopy[i] = arrCopy[i], arrCopy[c[i]]
				}
				c[i] += 1
				i = 0
				return arrCopy, hasMorePermutations
			}
			c[i] = 0
			i++
		}
		return []string{}, false
	}
}

func differByOneCharacter(s1, s2 string) bool {
	difference := false
	for i := range s1 {
		if s1[i] != s2[i] {
			if difference {
				return false
			}
			difference = true
		}
	}
	return difference
}

// Challenge
// https://app.codesignal.com/arcade/intro/level-7/PTWhv2oWqd6p4AHB9
func stringsRearrangement(inputArray []string) bool {
	sort.Strings(inputArray)
	generatePermutation := newPermutationsGenerator(inputArray)

	hasNext := true
permutationLoop:
	for hasNext {
		permutation, next := generatePermutation()
		hasNext = next
		for i := 1; i < len(permutation); i++ {
			if !differByOneCharacter(permutation[i-1], permutation[i]) {
				continue permutationLoop
			}
		}
		return true
	}

	return false
}

func main() {
	slice := []string{"abc", "abx", "axx", "abc"}
	fmt.Println(stringsRearrangement(slice)) //Should return false
}
