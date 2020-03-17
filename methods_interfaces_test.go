package gone

import (
	"fmt"
	"math"
	"testing"
)

func TestMethodsAndInterfaces(t *testing.T)  {
	v := Vertex{3, 4}
	fmt.Println(v.AbsMethod())
	fmt.Println(v.Scale(5))
	v.ScaleWOPointer(2)
	fmt.Println(v)
	fmt.Println(Scale(&v, 2))
	fmt.Println(Abs(Vertex{3, 4}))
	fmt.Println(Abs(Vertex{3, 4}))
	fmt.Println(MyFloat(-math.Sqrt2))
}

/* A METHOD is just a function with a receiver argument */
func (v Vertex) AbsMethod() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

/* a fluent METHOD with a pointer receiver */
func (v *Vertex) Scale(f float64) Vertex {
	v.X = v.X * f
	v.Y = v.Y * f
	return *v
}

/* a fluent method WITHOUT a pointer receiver copies over */
func (v Vertex) ScaleWOPointer(f float64) Vertex {
	v.X = v.X * f
	v.Y = v.Y * f
	return v
}

/*
	A function with a pointer
	Use a pointer to avoid copying the value on each method call.
	This if the receiver is a large struct, for example.
*/
func Scale(v *Vertex, f float64) Vertex {
	v.X = v.X * f
	v.Y = v.Y * f
	return *v
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}


