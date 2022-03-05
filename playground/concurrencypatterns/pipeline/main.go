package main

import "fmt"

func generate(from, to int) <-chan int {
	out := make(chan int, to-from-1)
	for i := from; i < to; i++ {
		out <- i
	}
	close(out)

	return out
}

func square(source <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for value := range source {
			out <- value * value
		}
		close(out)
	}()

	return out
}

func addOne(source <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for value := range source {
			out <- value + 1
		}
		close(out)
	}()

	return out
}

func main() {
	for value := range addOne(square(generate(1, 10))) {
		fmt.Println(value)
	}
}
