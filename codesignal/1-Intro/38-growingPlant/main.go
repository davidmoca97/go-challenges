package main

import (
	"fmt"
	"math"
)

func growingPlant(upSpeed int, downSpeed int, desiredHeight int) int {
	// y = desiredHeight
	// a = upSpeed
	// b = downSpeed
	// x = days

	// y = a + (a-b)(x-1)
	// x = ((y - a) / (a - b)) + 1

	if desiredHeight < upSpeed {
		return 1
	}

	return int(math.Ceil((float64(desiredHeight-upSpeed) / float64(upSpeed-downSpeed)) + 1))
}

// Challenge:
// https://app.codesignal.com/arcade/intro/level-9/xHvruDnQCx7mYom3T
func main() {
	fmt.Println(growingPlant(100, 10, 910)) // Should print 10
}
