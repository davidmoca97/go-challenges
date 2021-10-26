package main

import (
	"fmt"
	"strings"
)

/*
 * Complete the 'staircase' function below.
 *
 * The function accepts INTEGER n as parameter.
 */
func staircase(n int32) {
	for i := 0; i < int(n); i++ {
		template := fmt.Sprintf("%%%vv\n", n)
		fmt.Printf(template, strings.Repeat("#", i+1))
	}
}

// Challenge:
// https://www.hackerrank.com/challenges/staircase/problem
func main() {
	/*
		Should print:
			   #
			  ##
			 ###
			####
	*/
	staircase(4)
}
