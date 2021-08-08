package main

import "fmt"

func knapsackLight(value1 int, weight1 int, value2 int, weight2 int, maxW int) int {
	if weight1+weight2 <= maxW {
		return value1 + value2
	}
	if weight1 <= maxW && weight2 <= maxW {
		if value1 > value2 {
			return value1
		}
		return value2
	}
	if weight1 <= maxW {
		return value1
	}
	if weight2 <= maxW {
		return value2
	}
	return 0
}

// Challenge:
// https://app.codesignal.com/arcade/intro/level-9/r9azLYp2BDZPyzaG2
func main() {
	fmt.Println(knapsackLight(10, 5, 6, 4, 8)) //should print 10
}
