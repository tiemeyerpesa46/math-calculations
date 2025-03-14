package main

import "fmt"

func main() {
	// Add two numbers together
	result := addition(4, 5)
	fmt.Println("The result of 4 + 5 =", result)

	// Subtract one number from another
	result = subtraction(7, 3)
	fmt.Println("The result of 7 - 3 =", result)

	// Multiply two numbers together
	result = multiplication(4, 5)
	fmt.Println("The result of 4 x 5 =", result)

	// Divide one number by another
	result = division(10, 2)
	fmt.Println("The result of 10 / 2 =", result)
}

func addition(a int, b int) int {
	return a + b
}

func subtraction(a int, b int) int {
	return a - b
}

func multiplication(a int, b int) int {
	return a * b
}

func division(a int, b int) int {
	if b == 0 {
		return math.MaxInt32
	}
	return a / b
}
