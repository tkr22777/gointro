package gone

import (
	 "rsc.io/quote/v3"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func Gone() string {
	return "Hello, world!"
}

func Quote() string  {
	return quote.HelloV3()
}

func Proverb() string {
	return quote.Concurrency()
}

