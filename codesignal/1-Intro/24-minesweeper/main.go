package main

import "fmt"

func boolToInt(v bool) int {
	if v {
		return 1
	}
	return 0
}

func getCellValue(matrix [][]bool, i, j int) int {
	top := i > 0
	right := j < len(matrix[i])-1
	bottom := i < len(matrix)-1
	left := j > 0
	topRight := top && right
	topLeft := top && left
	bottomRight := bottom && right
	bottomLeft := bottom && left
	var counter int

	if top {
		counter += boolToInt(matrix[i-1][j])
	}
	if right {
		counter += boolToInt(matrix[i][j+1])
	}
	if bottom {
		counter += boolToInt(matrix[i+1][j])
	}
	if left {
		counter += boolToInt(matrix[i][j-1])
	}
	if topRight {
		counter += boolToInt(matrix[i-1][j+1])
	}
	if topLeft {
		counter += boolToInt(matrix[i-1][j-1])
	}
	if bottomRight {
		counter += boolToInt(matrix[i+1][j+1])
	}
	if bottomLeft {
		counter += boolToInt(matrix[i+1][j-1])
	}

	return counter
}

func minesweeper(matrix [][]bool) [][]int {
	result := make([][]int, len(matrix))
	for i, _ := range matrix {
		result[i] = make([]int, len(matrix[i]))
		for j, _ := range matrix[i] {
			result[i][j] = getCellValue(matrix, i, j)
		}
	}
	return result
}

// Challenge link
// https://app.codesignal.com/arcade/intro/level-5/ZMR5n7vJbexnLrgaM
func main() {
	input := [][]bool{
		{true, false, false},
		{false, true, false},
		{false, false, false},
	}
	fmt.Println(minesweeper(input)) // Should print [[1,2,1], [2,1,1], [1,1,1]
}
