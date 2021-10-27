package main

import "fmt"

func birthday(chocolateBar []int32, ronsBirthday int32, ronsBirthMonth int32) int32 {
	chocolateLen := int32(len(chocolateBar))
	var i int32
	var count int32
	for ; i <= chocolateLen-ronsBirthMonth; i++ {
		var sum int32
		for _, square := range chocolateBar[i : i+ronsBirthMonth] {
			sum += square
		}
		if sum == ronsBirthday {
			count++
		}
	}
	return count
}

// Challenge:
// https://www.hackerrank.com/challenges/the-birthday-bar
func main() {
	// Should print 2
	fmt.Println(birthday([]int32{2, 2, 1, 3, 2}, 4, 2))
}
