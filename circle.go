package golang_united_school_homework

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

const Pi = 3.14

func (c Circle) CalcPerimeter() float64 {
	p := 2 * Pi * c.Radius
	return p
}

func (c Circle) CalcArea() float64 {
	a := Pi * c.Radius * c.Radius
	return a
}
