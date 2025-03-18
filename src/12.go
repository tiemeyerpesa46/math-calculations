package main
import "math"
func add(x int, y int) int {
    return x + y
}
func subtract(x int, y int) int {
    return x - y
}
func multiply(x int, y int) int {
    return x * y
}
func divide(x int, y int) int {
    if y == 0 {
        panic("Cannot divide by zero")
    }
    return x / y
}
func main() {
    var x, y int
    fmt.Scan(&x)
    fmt.Scan(&y)
    result := add(x, y)
    fmt.Println("Result: ", result)
}