package main

import "fmt"

func main() {
	ch := make(chan string) // create a channel

	// Start a goroutine
	go func() {
		ch <- "Hello from Goroutine" // send value
	}()

	msg := <-ch // receive value
	fmt.Println(msg)
}
