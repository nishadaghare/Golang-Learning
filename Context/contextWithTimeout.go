package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	select {

	case <-time.After(3 * time.Second):
		fmt.Println("Task completed")

	case <-ctx.Done():
		fmt.Println("Timeout:", ctx.Err())
	}
}
