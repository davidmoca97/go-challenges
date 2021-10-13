package main

import (
	"fmt"
	"time"
)

func validTime(t string) bool {
	if _, err := time.Parse("15:04", t); err != nil {
		return false
	}
	return true
}

// Challenge:
// https://app.codesignal.com/arcade/intro/level-12/ywMyCTspqGXPWRZx5
func main() {
	fmt.Println(validTime("13:58")) // Should print true
	fmt.Println(validTime("25:51")) // Should print false
	fmt.Println(validTime("02:76")) // Should print false
}
