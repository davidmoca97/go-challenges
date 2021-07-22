package main

import "fmt"

func arrayReplace(inputArray []int, elemToReplace int, substitutionElem int) []int {
	for i := 0; i < len(inputArray); i++ {
		if inputArray[i] == elemToReplace {
			inputArray[i] = substitutionElem
		}
	}
	return inputArray
}

// Challenge link:
// https://app.codesignal.com/arcade/intro/level-6/mCkmbxdMsMTjBc3Bm
func main() {
	input := []int{1, 2, 1}
	fmt.Println(arrayReplace(input, 1, 3)) //Should print [3, 2, 3]
}
