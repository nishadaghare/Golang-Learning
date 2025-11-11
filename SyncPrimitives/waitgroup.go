package main

import (
	"fmt"
	"sync"
)

// When the main function finishes, all goroutines stop, even if they haven’t completed.
// To prevent this, use sync.WaitGroup
// A WaitGroup waits for a collection of goroutines to finish executing.
// add the number of goroutines to wait for using Add(),
// and each goroutine calls Done() when finished.
// The main goroutine calls Wait() to block until all are done.

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Worker", id, "done")
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("All workers done")
}
