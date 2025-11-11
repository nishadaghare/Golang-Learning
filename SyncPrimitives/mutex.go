package main

import (
	"fmt"
	"sync"
)

// A Mutex is used to protect shared data from being accessed by multiple goroutines at the same time.
// It allows only one goroutine to access a critical section at once.
func main() {
	var mu sync.Mutex
	counter := 0

	for i := 0; i < 3; i++ {
		go func() {
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	fmt.Scanln() // wait
	fmt.Println("Counter:", counter)
}
