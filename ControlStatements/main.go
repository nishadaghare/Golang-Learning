package main

import "fmt"

func main() {
	x := 10
	if x > 5 {
		fmt.Println("X is greater than 5")
	} else {
		fmt.Println("X is less than 5")
	}

	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("another day")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//range keyword is used to iterate through the elements of an array, slice or map
	arr := []string{"Banana", "Mango", "Apple"}

	for idx, value := range arr {
		fmt.Println(idx, value)
	}
}
