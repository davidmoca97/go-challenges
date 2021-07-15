package main

import "fmt"

func arrayChange(inputArray []int) int {
    var moves int
    for i:=1; i<len(inputArray); i++ {
        if inputArray[i-1] < inputArray[i] {
            continue
        }
        newMoves := inputArray[i-1] - inputArray[i] + 1
        inputArray[i] = inputArray[i] + newMoves
        moves += newMoves
    }
    return moves
}

// Challenge link
// https://app.codesignal.com/arcade/intro/level-4/xvkRbxYkdHdHNCKjg
func main() {
	slice := []int{-1000, 0, -2, 0}
	fmt.Println(arrayChange(slice)) //Should print 5
}
