package main

import "fmt"

func differentSquares(matrix [][]int) int {
	dictionary := make(map[string]bool)

	for i := 0; i < len(matrix)-1; i++ {
		for j := 0; j < len(matrix[0])-1; j++ {
			entry := fmt.Sprintf("%d,%d,%d,%d", matrix[i][j], matrix[i][j+1], matrix[i+1][j], matrix[i+1][j+1])
			dictionary[entry] = true
		}
	}

	return len(dictionary)
}

// Challenge:
// https://app.codesignal.com/arcade/intro/level-12/fQpfgxiY6aGiGHLtv
func main() {
	matrix := [][]int{
		{1, 2, 1},
		{2, 2, 2},
		{2, 2, 2},
		{1, 2, 3},
		{2, 2, 1},
	}
	fmt.Println(differentSquares(matrix)) //Should print 6
}
