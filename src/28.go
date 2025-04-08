package main

import "fmt"

func main() {
    sum := 1 + 2 * 3 - 4 / 5
    print(sum)
}

func print(x int) {
    fmt.Println("The result is:", x)
}
