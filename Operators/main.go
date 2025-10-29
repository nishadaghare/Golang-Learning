package main

import "fmt"

func main() {
	a, b := 10, 5
	fmt.Println(a+b, a-b, a*b, a/b, a%b)      //Arithmetic Operators
	fmt.Println(a == b, a != b, a < b, a > b) //Comparison Operators
	fmt.Println(a&b, a|b, a^b, a<<1, a>>1)    //Bitwise Operators
	x, y := true, false
	fmt.Println(x && y, x || y, !x) //Logical Operators
}
