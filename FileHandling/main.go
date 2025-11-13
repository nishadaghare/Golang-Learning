package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	filename := "example.txt"

	// Create a new file
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	fmt.Println("File created successfully!")

	// Write initial content
	_, err = file.WriteString("Hello, Golang!\nThis is a file handling example.\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		file.Close()
		return
	}
	file.Close() // close immediately
	fmt.Println("Initial content written!")

	//  Append more content
	file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file for append:", err)
		return
	}
	_, err = file.WriteString("Appended line: File handling is fun!\n")
	if err != nil {
		fmt.Println("Error appending to file:", err)
		file.Close()
		return
	}
	file.Close() //close before next operation
	fmt.Println("Content appended!")

	// Read entire file content using os.ReadFile
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("\n--- File Content (ReadFile) ---")
	fmt.Println(string(data))

	// Read line-by-line using bufio.Scanner
	file, err = os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file for reading:", err)
		return
	}
	fmt.Println("\n--- File Content (Line by Line) ---")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error:", err)
	}
	file.Close() // close after reading

	// Read using io.ReadAll
	file, err = os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading with io.ReadAll:", err)
		file.Close()
		return
	}
	file.Close() // close before delete
	fmt.Println("\n--- File Content (io.ReadAll) ---")
	fmt.Println(string(content))

	// Delete the file
	err = os.Remove(filename)
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}
	fmt.Println("\nFile deleted successfully!")
}
