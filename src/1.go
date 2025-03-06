package main

import "fmt"

func main() {
	a := 42
	b := 13
	result := addition(a, b)
	fmt.Println("Result:", result)
}

func addition(a, b int) int {
	return a + b
}
