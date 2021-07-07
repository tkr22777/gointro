package main

import (
	"fmt"
	"testing"
)

func TestStrings(t *testing.T) {
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + " World")
}
