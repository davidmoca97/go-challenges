package main

import "fmt"

/*
 * Complete the 'gradingStudents' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY grades as parameter.
 */

func gradingStudents(grades []int32) []int32 {
	roundedGrades := make([]int32, len(grades))

	for i := range grades {
		if grades[i] >= 38 && grades[i]%5 >= 3 {
			roundedGrades[i] = grades[i] + (5 - grades[i]%5)
			continue
		}
		roundedGrades[i] = grades[i]
	}

	return roundedGrades

}

// Challenge:
// https://www.hackerrank.com/challenges/grading/problem
func main() {
	res := gradingStudents([]int32{73, 67, 38, 33})
	fmt.Println(res)
}
