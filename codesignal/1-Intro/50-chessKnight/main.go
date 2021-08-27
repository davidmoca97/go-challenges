package main

import (
	"fmt"
	"strconv"
)

func getCoordenates(cell string) (x, y int) {
	x = int(cell[0] - 'a')
	y, _ = strconv.Atoi(string(cell[1]))
	return x, y - 1
}

func chessKnight(cell string) int {
	x, y := getCoordenates(cell)
	count := 0

	cases := map[string]bool{
		"case 1": y < 6 && x+1 <= 7,
		"case 2": y < 6 && x-1 >= 0,
		"case 3": y < 7 && x+2 <= 7,
		"case 4": y < 7 && x-2 >= 0,
		"case 5": y > 1 && x+1 <= 7,
		"case 6": y > 1 && x-1 >= 0,
		"case 7": y > 0 && x+2 <= 7,
		"case 8": y > 0 && x-2 >= 0,
	}

	for _, v := range cases {
		if v {
			count++
		}
	}

	return count
}

// Challenge:
// https://app.codesignal.com/arcade/intro/level-11/pwRLrkrNpnsbgMndb
func main() {
	fmt.Println(chessKnight("a1")) // Should print 2
	fmt.Println(chessKnight("c2")) // Should print 6
}
