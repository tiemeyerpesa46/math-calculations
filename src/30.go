package main

import (
	"fmt"
)

func calculate(a float64, b float64) float64 {
	return a + b
}

func main() {
	result := calculate(3.0, 5.0)
	fmt.Println("The result is:", result)
}
