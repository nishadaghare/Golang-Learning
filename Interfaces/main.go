package main

import "fmt"

type Shape interface {
	Area() float64
}

type Square struct {
	side float64
}

func (s Square) Area() float64 {
	return s.side * s.side
}

func main() {
	sq := Square{side: 4}

	var sh Shape = sq

	fmt.Println("Area of square:", sh.Area())
}
