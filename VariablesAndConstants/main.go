package main

import "fmt"

var message = "Global" //package-level scope
func main() {

	//Short Declaration
	x := 10
	name := "Go"
	fmt.Println(x, name)

	//Using var
	var country = "India"
	var isValid bool // zero value false
	fmt.Println(country, isValid)

	//var block
	var (
		a int
		b string
		c float64
	)
	fmt.Println(a, b, c)

	local := "Local"

	fmt.Println(message)
	fmt.Println(local)

	//Type Conversion
	var var1 int = 10
	var var2 float64 = float64(var1)
	fmt.Printf("Type of x: %T\n", var1)
	fmt.Printf("Type of y: %T\n", var2)

	var m float32 = 3.141592653589793
	var n float64 = 3.141592653589793

	fmt.Println(m) // 3.1415927  (less precision)
	fmt.Println(n) // 3.141592653589793 (more precise)

	//constant - cannot change once defined.
	const PI = 3.14
	// PI = 2.5 this will give error
	fmt.Println(PI)

}
