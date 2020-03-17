package gone

import (
	"fmt"
	"math"
	"testing"
)

func TestAbsAssignment(t *testing.T) {
	fmt.Println(assignment())
	IInterface()
	fmt.Println(stringerInterface())
}

type Abser interface {
	Abs() float64
}

func (v* Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func assignment() float64 {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4} //// a  implements Abser
	a = f // a MyFloat implements Abser
	a = &v //can assign Vertex pointer since it (implicitly) implements Abser
	//a = v //cannot assign Vertex and does not implement Abser
	return a.Abs()
}

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T (implicitly) implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func IInterface() {
	var i I = T{"Hello"}
	i.M()
}

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func stringerInterface() string {
	a := Person{"Zaphod Beeblebrox", 9001}
	return fmt.Sprint(a)
}

