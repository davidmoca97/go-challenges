package main

import "fmt"

/*
 eq1 = x(k1JumpRate) + k1StartPoint; eq2 = x(k2JumpRate) + k2StartPoint

 x(k1JumpRate) + k1StartPoint = x(k2JumpRate) + k2StartPoint ->
 x(k1JumpRate) - x(k2JumpRate) = k2StartPoint - k1StartPoint ->
 x(k1JumpRate - k2JumpRate) = k2StartPoint - k1StartPoint ->

 x = (k2StartPoint - k1StartPoint) / (k1JumpRate - k2JumpRate)

 If x is an integer, the kangoroos will intersect with each other, otherwise they won't
*/
func kangaroo(k1StartPoint int32, k1JumpRate int32, k2StartPoint int32, k2JumpRate int32) string {
	if k2JumpRate >= k1JumpRate {
		return "NO"
	}
	if (k2StartPoint-k1StartPoint)%(k1JumpRate-k2JumpRate) == 0 {
		return "YES"
	}
	return "NO"
}

func main() {
	fmt.Println(kangaroo(0, 3, 4, 2)) // Should print YES
	fmt.Println(kangaroo(0, 2, 5, 3)) // Should print NO
}
