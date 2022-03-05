package main

import (
	"crypto/md5"
	"fmt"
	"sync"
)

func generate(from, to int) <-chan int {
	out := make(chan int)
	go func() {
		for i := from; i < to; i++ {
			out <- i
		}
		close(out)
	}()

	return out
}

func square(done <-chan struct{}, source <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for value := range source {
			select {
			case out <- value * value:
			case <-done:
				break
			}
		}
	}()

	return out
}

func addOne(done <-chan struct{}, source <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for value := range source {
			select {
			case out <- value + 1:
			case <-done:
				break
			}
		}
	}()

	return out
}

func merge(done <-chan struct{}, channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, ch := range channels {
		go func(ch <-chan int) {
			defer wg.Done()
			for value := range ch {
				select {
				case out <- value:
				case <-done:
					return
				}
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	done := make(chan struct{})
	defer close(done)

	md5.New()
	c1 := square(done, generate(10, 20))
	c2 := square(done, generate(0, 10))

	for value := range merge(done, c1, c2) {
		fmt.Println(value)
	}
}
