package main

import (
	"fmt"
	"time"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */
func timeConversion(s string) string {
	time, err := time.Parse("03:04:05PM", s)
	if err != nil {
		panic("Time conversion failed")
	}
	return time.Format("15:04:05")
}

// Challenge:
// https://www.hackerrank.com/challenges/time-conversion/problem
func main() {
	fmt.Println(timeConversion("12:01:00AM"))
}
