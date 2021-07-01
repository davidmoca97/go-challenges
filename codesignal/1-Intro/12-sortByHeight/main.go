package main

import "sort"

func sortByHeight(a []int) []int {
	dupe := make([]int, len(a))
	copy(dupe, a)
	sort.Ints(dupe)

	dupeIdx := len(a) - 1
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == -1 {
			continue
		}
		a[i] = dupe[dupeIdx]
		dupeIdx--
	}

	return a
}

func main() {
	slice := []int{-1, 150, 190, 170, -1, -1, 160, 180}
	// Should return [-1, 150, 160, 170, -1, -1, 180, 190]
	sortByHeight(slice)
}
