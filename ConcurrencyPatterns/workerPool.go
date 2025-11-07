package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("Worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("Worker", id, "finished job", j)
		result <- j * 2
	}
}

func main() {
	numofjobs := 5
	jobs := make(chan int, numofjobs)
	result := make(chan int, numofjobs)

	for w := 0; w < 3; w++ {
		go worker(w, jobs, result)
	}

	for j := 0; j < numofjobs; j++ {
		jobs <- j
	}

	close(jobs)

	for i := 0; i < numofjobs; i++ {
		<-result
	}
	close(result)

}
