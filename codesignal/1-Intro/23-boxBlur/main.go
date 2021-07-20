package main

import "fmt"

func getAvg(matrix [][]int, startInColumn int) int {
	var total int
	for i := 0; i < 3; i++ {
		for j := startInColumn; j < startInColumn+3; j++ {
			total += matrix[i][j]
		}
	}
	return total / 9
}

func boxBlur(image [][]int) [][]int {
	var blurredImage [][]int

	for i := 1; i < len(image)-1; i++ {
		var row []int
		for j := 1; j < len(image[i])-1; j++ {
			avg := getAvg(image[i-1:i+2], j-1)
			row = append(row, avg)
		}
		blurredImage = append(blurredImage, row)
	}

	return blurredImage
}

// Challenge:
// https://app.codesignal.com/arcade/intro/level-5/5xPitc3yT3dqS7XkP
func main() {
	matrix := [][]int{
		{36, 0, 18, 9},
		{27, 54, 9, 0},
		{81, 63, 72, 45},
	}
	fmt.Println(boxBlur(matrix)) // Should print [[40,30]]
}
