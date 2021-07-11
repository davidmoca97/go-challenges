package main

import (
	"fmt"
	"strings"
)

func printPicture(picture []string) {
	for _, row := range picture {
		fmt.Println(row)
	}
}

func addBorder(picture []string) []string {
	horizontalBorder := strings.Repeat("*", len(picture[0])+2)
	pictureWithBorder := make([]string, len(picture)+2)
	pictureWithBorder[0] = horizontalBorder
	idx := 1
	for _, row := range picture {
		pictureWithBorder[idx] = "*" + row + "*"
		idx++
	}
	pictureWithBorder[len(pictureWithBorder)-1] = horizontalBorder
	return pictureWithBorder
}

// Challenge link
// https://app.codesignal.com/arcade/intro/level-4/ZCD7NQnED724bJtjN/solutions
func main() {
	picture := []string{
		"abc",
		"ded",
	}
	/* Should return
	["*****",
	"*abc*",
	"*ded*",
	"*****"]
	*/
	pictureWithBorder := addBorder(picture)
	printPicture(pictureWithBorder)
}
