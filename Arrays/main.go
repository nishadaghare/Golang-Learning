package main

import "fmt"

func main() {
	var numbers [5]int // declared but not initialized → all zeros initially

	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	fmt.Println(numbers)

	var arr = [3]string{"Apple", "Banana", "Watermelon"} //Declaration with Initialization
	fmt.Println(arr)
	arr[2] = "Mango" // Modify
	fmt.Println(arr)

	var arr1 = []string{"Hello", "Nishada", "Good Morning !"}
	for idx, val := range arr1 { //print using range
		fmt.Println(idx, val)
	}

	odd := [...]int{1, 3, 5, 7, 9} //Go figure out the size
	fmt.Println(odd)

	arr2 := []float32{1.5, 5.6, 15.3}

	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

	x := [2]int{1, 2}
	y := [2]int{1, 2}
	fmt.Println(x == y)

	a := [3]int{1, 2, 3}
	b := a // b is a copy of a
	b[0] = 100

	fmt.Println(a)
	fmt.Println(b)
}
