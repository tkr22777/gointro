package simplepackage

import (
	"fmt"
	"math"
	"testing"
)

func TestAbsAssignment(t *testing.T) {
	fmt.Println(variable_assignment())
	IInterface()
	fmt.Println(stringerInterface())
}

type Abser interface {
	Abs() float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func variable_assignment() float64 {
	var a Abser               // a implements Abser
	f := MyFloat(-math.Sqrt2) // MyFloat implements Abser
	v := Vertex{3, 4}         // Vertex (implicitly) implements Abser when it is pointer?
	a = f                     // Can be assigned to a (Abser type) var
	a = &v                    // can assign Vertex pointer since it, the pointer, implicitly implements Abser
	//a = v 		  // but cannot assign Vertex and as it does not implement Abser
	return a.Abs()
}

type Worker interface {
	work()
}

type Engineer struct {
	feild string
}

// This method means type T (implicitly) implements the interface I,
// but we don't need to explicitly declare that it does so.
func (e Engineer) work() {
	fmt.Println("Wow wow wee wa! Imma open the dashboard first!")
}

func IInterface() {
	var allDayWorker Worker = Engineer{"Software"}
	allDayWorker.work()
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func stringerInterface() string {
	a := Person{"Zaphod Beeblebrox", 9001}
	return fmt.Sprint(a)
}
