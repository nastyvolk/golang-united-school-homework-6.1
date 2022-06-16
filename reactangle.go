package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

func (r Rectangle) CalcPerimeter() float64 {
	p := (r.Height + r.Weight) * 2
	return p
}

func (r Rectangle) CalcArea() float64 {
	a := r.Height * r.Weight
	return a
}
