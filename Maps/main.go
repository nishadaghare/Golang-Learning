package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Apple"] = 10
	m["banana"] = 5
	m["Apple"] = 15 //update
	fmt.Println(m)
	delete(m, "banana")
	fmt.Println(m)

	map1 := map[string]int{"Go": 1, "Python": 2}
	for k, v := range map1 {
		fmt.Println(k, v)
	}

	val, ok := map1["Python"]
	fmt.Println("Value:", val, "Exists:", ok)

}
