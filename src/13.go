package main

import "fmt"

func main() {
	// Generate a random number between 1 and 100
	num1 := rand.Intn(100) + 1
	// Generate another random number between 1 and 100
	num2 := rand.Intn(100) + 1

	// Calculate the sum of the two numbers
	sum := num1 + num2

	fmt.Println("The sum is:", sum)
}