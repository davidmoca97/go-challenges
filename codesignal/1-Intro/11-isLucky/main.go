package main

import "strconv"

func isLucky(n int) bool {
	str := strconv.Itoa(n)
	count := 0

	for i := 0; i < len(str)/2; i++ {
		digit, err := strconv.Atoi(string(str[i]))
		if err != nil {
			return false
		}
		count += digit
	}
	for i := len(str) / 2; i < len(str); i++ {
		digit, err := strconv.Atoi(string(str[i]))
		if err != nil {
			return false
		}
		count -= digit
	}

	return count == 0
}

func main() {
	// should return true
	isLucky(1230)
	// Should return false
	isLucky(239017)
}
