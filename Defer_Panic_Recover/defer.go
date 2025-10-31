package main

import "fmt"

func main() {
	//In Go, defer is a keyword used to delay the execution of a function until the surrounding function finishes.
	fmt.Println("Start")

	defer fmt.Println("Deferred statement") // This runs at the end of main()

	fmt.Println("End")

}
