package main

import "fmt"

func sudoku(grid [][]int) bool {

	for i := range grid {
		rowCount, colCount := 0, 0
		for j := range grid[i] {
			rowCount += grid[i][j]
			colCount += grid[j][i]
		}
		// All columns and rows should add up 45
		// 1 + 2 + 3 + ... + 9 = 45
		if rowCount != 45 || colCount != 45 {
			return false
		}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			dictionary := make(map[int]bool)

			for k := i; k < i+3; k++ {
				for l := j; l < j+3; l++ {
					dictionary[grid[k][l]] = true
				}
			}

			// Make sure all of the 3x3 quadrants of the sudoku
			// include all of the 9 digits [1-9]
			if len(dictionary) != 9 {
				return false
			}
		}
	}

	return true

}

// Challenge:
// https://app.codesignal.com/arcade/intro/level-12/tQgasP8b62JBeirMS
func main() {
	test := [][]int{
		{1, 3, 2, 5, 4, 6, 9, 8, 7},
		{4, 6, 5, 8, 7, 9, 3, 2, 1},
		{7, 9, 8, 2, 1, 3, 6, 5, 4},
		{9, 2, 1, 4, 3, 5, 8, 7, 6},
		{3, 5, 4, 7, 6, 8, 2, 1, 9},
		{6, 8, 7, 1, 9, 2, 5, 4, 3},
		{5, 7, 6, 9, 8, 1, 4, 3, 2},
		{2, 4, 3, 6, 5, 7, 1, 9, 8},
		{8, 1, 9, 3, 2, 4, 7, 6, 5},
	}
	fmt.Println(sudoku(test)) // Should print true
}
