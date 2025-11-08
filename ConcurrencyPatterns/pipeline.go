package main

import (
	"fmt"
)

// The Pipeline concurrency pattern in Go involves a series of stages, where the output of one stage serves as the input for the next.
// generator: sends numbers into a channel
func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// square: receives numbers, squares them, and sends to next stage
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// double: doubles the squared numbers
func double(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * 2
		}
		close(out)
	}()
	return out
}

func main() {
	// Build the pipeline
	nums := generator(1, 2, 3, 4, 5)
	squares := square(nums)
	doubles := double(squares)

	// Consume the output
	for result := range doubles {
		fmt.Println(result)
	}
}
