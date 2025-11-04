package main

import "fmt"

type Counter struct {
	Count int
}

func (c Counter) IncrementVal() { // value receiver
	c.Count++
}

func (c *Counter) IncrementPtr() { // pointer receiver
	c.Count++
}

func main() {
	c := Counter{5}
	c.IncrementVal()
	fmt.Println("After value increment:", c.Count) // unchanged

	c.IncrementPtr()
	fmt.Println("After pointer increment:", c.Count) // updated
}
