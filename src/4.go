func add(a int, b int) int {
    return a + b
}
func subtract(a int, b int) int {
    return a - b
}
func multiply(a int, b int) int {
    return a * b
}
func divide(a int, b int) int {
    if b == 0 {
        panic("division by zero")
    }
    return a / b
}