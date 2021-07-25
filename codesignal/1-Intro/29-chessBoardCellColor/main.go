package main

import (
	"fmt"
	"strconv"
)

func getCoordinates(cell string) (x, y int) {
	letterToNumber := map[byte]int{
		'A': 0,
		'B': 1,
		'C': 2,
		'D': 3,
		'E': 4,
		'F': 5,
		'G': 6,
		'H': 7,
	}
	yIndex, _ := strconv.Atoi(string(cell[1]))
	return letterToNumber[cell[0]], yIndex - 1
}

func chessBoardCellColor(cell1 string, cell2 string) bool {
	smallBoard := [][]bool{
		{false, true},
		{true, false},
	}
	x1, y1 := getCoordinates(cell1)
	x2, y2 := getCoordinates(cell2)

	color1 := smallBoard[y1%2][x1%2]
	color2 := smallBoard[y2%2][x2%2]

	return color1 == color2
}

// Challenge:
// https://app.codesignal.com/arcade/intro/level-6/t97bpjfrMDZH8GJhi
func main() {
	fmt.Println(chessBoardCellColor("A1", "C3")) //Should return true
}
