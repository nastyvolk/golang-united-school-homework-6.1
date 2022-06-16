package golang_united_school_homework

import "math"

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (c Circle) CalcPerimeter() float64 {
	p := 2 * math.Pi * c.Radius
	return p
}

func (c Circle) CalcArea() float64 {
	a := math.Pi * c.Radius * c.Radius
	return a
}
