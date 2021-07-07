package simplepackage

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func TestPkgVarFunc(t *testing.T) {
	fmt.Println("Start of test related to package, var and function:")
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
	fmt.Println(typeConversion())
	fmt.Println(defaultTypeInference())
	fmt.Println(constants())
}

func add(x int, y int) int {
	return x + y
}

func swap(first string, second string) (string, string) {
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
	var unByte byte = 37    //uint8
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

func typeConversion() string {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	return fmt.Sprint(x, y, z)
}

func defaultTypeInference() string {
	i := 294967292    // int32
	f := 3.142        //
	g := 0.867 + 0.5i //complex128
	return fmt.Sprintf("294967292 is of type:%T, 3.142 is of type:%T and 0.867 + 0.5i is of type:%T", i, f, g)
}

func constants() string {
	const PI = 3.1416
	const MORTAL = true
	return fmt.Sprint(PI, MORTAL)
}
