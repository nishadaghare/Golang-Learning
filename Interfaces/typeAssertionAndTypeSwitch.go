package main

import "fmt"

func main() {
	var v interface{} = 10

	// Type assertion
	if val, ok := v.(int); ok {
		fmt.Println("Int value:", val)
	}

	// Type switch
	switch t := v.(type) {
	case int:
		fmt.Println("Int:", t)
	case string:
		fmt.Println("String:", t)
	default:
		fmt.Println("Unknown type")
	}
}
