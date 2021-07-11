package main

import "fmt"

func alternatingSums(a []int) []int {
	sums := make([]int, 2)
	for idx, weight := range a {
		if (idx+1)%2 == 0 {
			sums[1] += weight
			continue
		}
		sums[0] += weight
	}
	return sums
}

// Link of the challenge
// https://app.codesignal.com/arcade/intro/level-4/cC5QuL9fqvZjXJsW9/solutions
func main() {
	nums := []int{50, 60, 60, 45, 70}
	res := alternatingSums(nums) // Should return [180, 105]
	fmt.Println(res)
}
