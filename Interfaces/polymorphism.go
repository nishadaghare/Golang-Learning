package main

import "fmt"

type Shape interface {
	Area() float64
}

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

func (s Square) Area() float64 {
	return s.side * s.side
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func PrintArea(sh Shape) {
	fmt.Println("Area:", sh.Area())
}

func main() {
	sq := Square{side: 4}
	cr := Circle{radius: 3}

	PrintArea(sq) // Area: 16
	PrintArea(cr) // Area: 28.26
}
