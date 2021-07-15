package main

import "fmt"

func arrayMaximalAdjacentDifference(inputArray []int) int {
    max := 0
    for i:=1; i<len(inputArray); i++ {
        difference := inputArray[i-1] - inputArray[i]
        if difference < 0 {
            difference *= -1
        }
        if difference > max {
            max = difference
        }
    }
    
    return max
}

// Challenge link
// https://app.codesignal.com/arcade/intro/level-5/EEJxjQ7oo7C5wAGjE
func main() {
	input := []int{-1, 4, 10, 3, -2}
	fmt.Println(arrayMaximalAdjacentDifference(input)) // should print 7
}
