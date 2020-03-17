package gone

import (
	"fmt"
	"testing"
)

type Vertex struct {
	X, Y float64
}

func TestStructSlicesMap(t *testing.T) {
	fmt.Println("Start of test related to struct, slices and map:")
	fmt.Println(pointers())
	fmt.Println(testStruct())
	fmt.Println(testStructPointer())
	fmt.Println(structLiteral())
	fmt.Println(array())
	fmt.Println(slices())
	fmt.Println(sliceRefers())
	fmt.Println(sliceLiteral())
	fmt.Println(sliceDefaults())
	fmt.Println(sliceLenAndCap())
	fmt.Println(nilSlice())
	fmt.Println(sliceWithMake())
	fmt.Println(appendToASlice())
	fmt.Println(sliceRange())
	fmt.Println(rangeSkipInderOrValue())
	fmt.Println(mapBegins())
	fmt.Println(mapLiterals())
}

func pointers() string {
	i, j := 2, 27
	p := &i //point to i
	q := &j
	*q = *q / 3 //divide j through the pointer
	return fmt.Sprintf("*p: %d, i: %d, j: %d", *p, i, j)
}

func testStruct() string {
	v := Vertex{1, 2}
	return fmt.Sprint("v:", v, " v.x:", v.X)
}

func testStructPointer() string {
	v := Vertex{1, 3}
	p := &v
	return fmt.Sprint("v:", v, " v.x (with p):", p.X)
}

func structLiteral() string {
	v1 := Vertex{1, 2}
	v2 := Vertex{X: 1}
	v3 := Vertex{}
	p := &Vertex{1, 2}
	return fmt.Sprint(v1, v2, v3, p)
}

func array() string {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	primes := [6]int{2, 3, 5, 7, 11, 13}
	return fmt.Sprint(a[0], a[1], primes)
}

func slices() string {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	return fmt.Sprint(s)
}

func sliceRefers() string {
	names := [4]string {
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	a := names[0:2]
	b := names[1:3]
	b[0] = "XXX"
	return fmt.Sprint(a, b, names)
}

func sliceLiteral() string {
	q := []int{2, 3, 5, 7, 11, 13}
	r := []bool{true, false, true, true, false, true}
	s := []struct{
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	return fmt.Sprint(q, r, s)
}

func sliceDefaults() string {
	s := []int{2, 3, 5, 7, 11, 13}
	p := s[1:4]
	s = p[:2]
	return fmt.Sprint(p, s)
}

func sliceLenAndCap() string {
	s := []int{2, 3, 5, 7, 11, 13}
	s = s[:0]
	p := s[:4]
	return fmt.Sprintf("S: len=%d , cap=%d %v, P: len=%d , cap=%d %v",
		len(s), cap(s), s, len(p), cap(p), p)
}

func nilSlice() string {
	var s []int
	isNil := false
	if s == nil {
		isNil = true
	}
	return fmt.Sprint(s, len(s), cap(s), isNil)
}

func sliceWithMake() string {
	s := make([]int, 5)
	p := make([]int, 0, 5)
	return fmt.Sprintf("S: len=%d , cap=%d %v, P: len=%d , cap=%d %v",
		len(s), cap(s), s, len(p), cap(p), p)
}

func appendToASlice() string {
	var s []int
	s = append(s, 0)
	s = append(s, 1)
	s = append(s, 2, 3, 4)
	return fmt.Sprintf("S: len=%d , cap=%d %v", len(s), cap(s), s)
}

func sliceRange() string {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	out := ""
	for i, v := range pow {
		out = out + fmt.Sprintf("2**%d = %d, ", i, v)
	}
	return out
}

func rangeSkipInderOrValue() string {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	out := "Vals:"
	for _, v := range pow {
		out = out + fmt.Sprintf("%d, ", v)
	}
	return out
}

func mapBegins() string {
	var m map[string]Vertex
	m = make(map[string]Vertex)
	m["Origin"] = Vertex{
		0, 0,
	}
	return fmt.Sprint(m["Origin"])
}

func mapLiterals() string {
	var m = map[string]Vertex{
		"Origin": {0, 0},
		"Top": {0, 100},
	}
	return fmt.Sprint(m)
}


