package main

import (
	"fmt"
	"time"
)

// simulate fetch
func fetchURL(id int, url string, results chan<- string) {
	fmt.Printf("Worker %d fetching %s\n", id, url)
	time.Sleep(time.Second) // simulate delay
	results <- fmt.Sprintf("Worker %d done with %s", id, url)
}

func main() {
	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://openai.com",
		"https://github.com",
	}

	results := make(chan string, len(urls))

	// --- Fan-Out: multiple workers fetching concurrently ---
	for i, url := range urls {
		go fetchURL(i+1, url, results)
	}

	// --- Fan-In: collect all results into one channel ---
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-results)
	}
}
