package main

import (
	"fmt"
)

func main() {
	var a int = 5
	var b int = 3

	sum := a + b
	product := a * b
	dividend := 10 / (a - b)
	
	fmt.Println("Sum:", sum)
	fmt.Println("Product:", product)
	fmt.Println("Dividend:", dividend)
}
