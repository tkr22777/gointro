package gone

import (
	"fmt"
	"math"
	"testing"
)

func TestFlowControl(t *testing.T) {
	fmt.Println("Start of test related to flow control:")
	fmt.Println( "Sum:", sum(10))
	fmt.Println( "Sqrt:", sqrt(-10))
}

func sum(num int) int {
	sum := 0
	for i:= 1; i <= num; i++ {
		sum += i
	}
	return sum
}

func sqrt(num int) string {
	if num < 0 {
		return sqrt(-num) + "i"
	}
	return fmt.Sprint(math.Sqrt(float64(num)))
}



