package gone

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func TestMath(t *testing.T) {
	fmt.Println("A random number:", rand.Intn(10))
	fmt.Println("Sqrt of 7:", math.Sqrt(7))
	fmt.Println("Pi:", math.Pi)
	fmt.Println("5 + 7 ->", add(5, 7))
	fmt.Println(swap("first", "second"))
}

func add(x int, y int) int {
	return x + y
}

func swap(first string, second string) (string, string)  {
	return second, first
}

