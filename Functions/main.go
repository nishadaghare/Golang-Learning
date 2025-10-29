package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func divide(a, b int) (int, int) { //Multiple Return Values
	return a / b, a % b
}

func rectangle(length, width int) (area int, perimeter int) { // Named Return Values
	area = length * width
	perimeter = 2 * (length + width)
	return // implicit
}

//Variadic Functions - You can pass any number of arguments using ...
func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

func main() {

	result := add(5, 3)
	fmt.Println(result)

	q, r := divide(10, 3)
	fmt.Println("Quotient:", q, "Remainder:", r)

	a, p := rectangle(2, 4)
	fmt.Println(a, p)

	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(1, 2))

	// Anonymous Function
	greet := func(name string) {
		fmt.Println("Hello", name)
	}
	greet("Nishada")

	// Immediately invoked function
	func(msg string) {
		fmt.Println(msg)
	}("Welcome to Go!")
}
