package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = append(slice, 6, 7, 8)

	fmt.Println(slice)

	var colors = []string{"Blue", "Red", "Pink"}

	for id, val := range colors {
		fmt.Println(id, val)
	}

	var colors1 = make([]string, 5)
	copy(colors1, colors)
	fmt.Println("Copied slice:", colors1)

	var s = [4]int{2, 4, 6, 8}
	sub := s[1:3]
	fmt.Println("Sliced (1:4):", sub)

}
