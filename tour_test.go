package gone

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func TestEtc(t *testing.T) {
	fmt.Println("A random number:", rand.Intn(10))
	fmt.Println("Sqrt of 7:", math.Sqrt(7))
	fmt.Println("Pi:", math.Pi)
	fmt.Println("5 + 7 ->", add(5, 7))
	fmt.Println(swap("first", "second"))
	fmt.Println(split(37))
	fmt.Println(varDeclaration())
	fmt.Println(varDeclaration2())
	fmt.Println(varDeclaration3())
	fmt.Println(varTypes())
	fmt.Println(zeroValues())
}

func add(x int, y int) int {
	return x + y
}

func swap(first string, second string) (string, string)  {
	return second, first
}

func split(doubleDigit int) (x, y int) {
	y = doubleDigit % 10
	x = doubleDigit / 10
	return
}

func varDeclaration() string {
	var c, python, java bool
	var i int
	return fmt.Sprint(i, c, python, java)
}

func varDeclaration2() string {
	var c, python, java = true, false, " no!"
	var i, j int = 1, 2
	return fmt.Sprint(i, j, c, python, java)
}

func varDeclaration3() string {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, " no!"
	return fmt.Sprint(i, j, k, c, python, java)
}

func varTypes() string {
	var isDead bool = false
	var firstName string = "Hey Jude!"
	var integerDe64 int64 = -3412
	var nonNegative uint64 = 12319
	var unByte byte = 37 //uint8
	var unRune rune = 31239 //uint32 or an unicode
	var fraction float64 = 93747.1319
	return fmt.Sprint(isDead, firstName, integerDe64, nonNegative, unByte, unRune, fraction)
}

func zeroValues() string {
	var i int
	var f float64
	var b bool
	var s string
	return fmt.Sprint(i, f, b, s)
}




