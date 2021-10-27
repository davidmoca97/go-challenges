package main

import "fmt"

/*
 * Complete the 'countApplesAndOranges' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER s - house start point
 *  2. INTEGER t - house finish point
 *  3. INTEGER a - apple tree
 *  4. INTEGER b - orange tree
 *  5. INTEGER_ARRAY apples
 *  6. INTEGER_ARRAY oranges
 */

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	var applesCount, orangesCount int32
	for _, apple := range apples {
		if apple+a >= s && apple+a <= t {
			applesCount++
		}
	}
	for _, orange := range oranges {
		if orange+b >= s && orange+b <= t {
			orangesCount++
		}
	}
	fmt.Println(applesCount)
	fmt.Println(orangesCount)
}

// Challenge:
// https://www.hackerrank.com/challenges/apple-and-orange/problem
func main() {
	// Should print:
	// 1
	// 2
	countApplesAndOranges(7, 10, 4, 12, []int32{2, 3, -4}, []int32{3, -2, -4})
}
