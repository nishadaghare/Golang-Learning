package main

import "fmt"

func printValue(v interface{}) {
	fmt.Println(v)
}

func main() {
	printValue(10)
	printValue("GoLang")
	printValue(3.14)
}
