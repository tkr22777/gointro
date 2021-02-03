package main

import (
    "fmt"
    "math"
    "math/rand"
)

func add(x int, y int) {
    return x + y
}

func main() {
    fmt.Println("My favorite number is", rand.Intn(10))
    fmt.Println(math.Pi)
}

