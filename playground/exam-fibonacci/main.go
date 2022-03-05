package main

import (
	"fmt"
	"sort"
	"time"
)

/*
 * Complete the 'ModuloFibonacciSequence' function below.
 *
 * The function accepts following parameters:
 *  1. chan bool requestChan
 *  2. chan int resultChan
 */

func ModuloFibonacciSequence(requestChan chan bool, resultChan chan int) {
	var a, b int = 1, 2
	idx := 0
	ticker := time.NewTicker(time.Millisecond * 10)
	for range requestChan {
		if idx == 0 {
			<-ticker.C
			resultChan <- 1
		}
		if idx == 1 {
			<-ticker.C
			resultChan <- 2
		}
		if idx > 1 {
			a, b = b, a+b
			<-ticker.C
			resultChan <- b % 1000000000
		}
		idx++
	}
	ticker.Stop()
}

func main() {
	arr := []string{"abc", "ab", "zzzz", "abcde", "a", "abcd"}
	sort.Slice(arr, func(i, j int) bool {
		if len(arr[i]) == len(arr[j]) {
			return arr[i][0] < arr[j][0]
		}
		return arr[i] < arr[j]
	})
	fmt.Println(arr)
}
