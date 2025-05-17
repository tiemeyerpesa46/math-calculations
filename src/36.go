package main

import (
    "fmt"
)

func main() {
    var x, y int
    fmt.Print("Enter first number: ")
    fmt.Scan(&x)
    fmt.Print("Enter second number: ")
    fmt.Scan(&y)
    result := x + y
    fmt.Printf("The sum of %d and %d is: %d\n", x, y, result)
}
