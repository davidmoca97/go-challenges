package main

import "fmt"

func getNextDirection(d string) string {
	switch d {
	case "EAST":
		return "SOUTH"
	case "SOUTH":
		return "WEST"
	case "WEST":
		return "NORTH"
	default:
		return "EAST"
	}
}

func newGenerator(n int) func() (x, y int) {
	leftLimit, topLimit := 0, 1
	rightLimit, bottomLimit := n-1, n-1
	x, y := -1, 0
	direction := "EAST"

	return func() (int, int) {
		switch direction {
		case "EAST":
			{
				x++
				if x+1 > rightLimit {
					direction = getNextDirection(direction)
					rightLimit--
				}
			}
		case "SOUTH":
			{
				y++
				if y+1 > bottomLimit {
					direction = getNextDirection(direction)
					bottomLimit--
				}
			}
		case "WEST":
			{
				x--
				if x-1 < leftLimit {
					direction = getNextDirection(direction)
					leftLimit++
				}
			}
		case "NORTH":
			{
				y--
				if y-1 < topLimit {
					direction = getNextDirection(direction)
					topLimit++
				}
			}
		}

		return x, y
	}
}

func spiralNumbers(n int) [][]int {
	nextCell := newGenerator(n)
	result := make([][]int, n)

	for i := range result {
		result[i] = make([]int, n)
	}

	for i := 1; i <= n*n; i++ {
		x, y := nextCell()
		result[y][x] = i
	}

	return result
}

// Challenge:
// https://app.codesignal.com/arcade/intro/level-12/uRWu6K8E7uLi3Qtvx/
func main() {
	fmt.Println(spiralNumbers(3)) // Should print [[1, 2, 3], [8, 9, 4], [7, 6, 5]]
}
