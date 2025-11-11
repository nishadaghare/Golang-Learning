package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Go Web Developer!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Running on :8080")
	http.ListenAndServe(":8080", nil)
}
