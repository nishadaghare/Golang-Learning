package main

import "fmt"

func updateValueByReferance(val *int) {
	*val = 100 // dereference and modify original value
}
func updateValue(val int) {
	val = 100
}

func main() {
	x := 24
	p := &x
	fmt.Println("address", p)
	fmt.Println("value", *p)

	num := 10
	updateValue(num)
	fmt.Println(num) // Output: 10 (unchanged)  -> Passing by value - orignal value is not changed

	num1 := 10
	updateValueByReferance(&num1)
	fmt.Println(num1) // Output: 100 (changed!)  -> Passing by Reference - orignal value is changed
}
