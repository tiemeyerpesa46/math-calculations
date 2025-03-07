
package main
import "fmt"
func main() {
	//Addition
	fmt.Println("1 + 2 = ", add(1, 2))
	fmt.Println("3 + 4 = ", add(3, 4))
	fmt.Println("5 + 6 = ", add(5, 6))

	//Subtraction
	fmt.Println("8 - 5 = ", subtract(8, 5))
	fmt.Println("10 - 3 = ", subtract(10, 3))
	fmt.Println("12 - 7 = ", subtract(12, 7))

	//Multiplication
	fmt.Println("2 * 4 = ", multiply(2, 4))
	fmt.Println("6 * 8 = ", multiply(6, 8))
	fmt.Println("10 * 5 = ", multiply(10, 5))

	//Division
	fmt.Println("24 / 4 = ", divide(24, 4))
	fmt.Println("32 / 2 = ", divide(32, 2))
	fmt.Println("64 / 8 = ", divide(64, 8))
}
func add(a, b int) int {
	return a + b
}
func subtract(a, b int) int {
	return a - b
}
func multiply(a, b int) int {
	return a * b
}
func divide(a, b int) int {
	return a / b
}