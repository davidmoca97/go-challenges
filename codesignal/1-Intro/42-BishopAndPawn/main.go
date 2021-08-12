package main

import (
	"fmt"
	"strconv"
)

func getCoordinates(str string) (x, y int) {
	letterToNumberMap := map[byte]int{
		'a': 0,
		'b': 1,
		'c': 2,
		'd': 3,
		'e': 4,
		'f': 5,
		'g': 6,
		'h': 7,
	}
	x = letterToNumberMap[str[0]]
	y, _ = strconv.Atoi(string(str[1]))
	return x, y
}

func bishopAndPawn(bishop string, pawn string) bool {
	bishopX, bishopY := getCoordinates(bishop)
	pawnX, pawnY := getCoordinates(pawn)

	return bishopX-bishopY == pawnX-pawnY || bishopX+bishopY == pawnX+pawnY
}

// Challenge:
// https://app.codesignal.com/arcade/intro/level-9/6M57rMTFB9MeDeSWo/
func main() {
	fmt.Println(bishopAndPawn("a1", "c3")) //Should print true
}
