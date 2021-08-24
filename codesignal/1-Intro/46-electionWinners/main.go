package main

import "fmt"

func electionsWinners(votes []int, k int) int {
	max := -1
	counter := 0
	multipleMax := false
	for _, vote := range votes {
		if vote > max {
			max = vote
			continue
		}
		if vote == max {
			multipleMax = true
		}
	}

	for _, vote := range votes {
		if vote == max && k == 0 && !multipleMax {
			counter++
		}
		if vote+k > max {
			counter++
		}
	}

	return counter
}

// Challenge:
// https://app.codesignal.com/arcade/intro/level-10/8RiRRM3yvbuAd3MNg
func main() {
	votes := []int{2, 3, 5, 2}
	k := 3
	fmt.Println(electionsWinners(votes, k)) //Should print 2
}
