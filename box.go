package golang_united_school_homework

import "errors"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) < b.shapesCapacity {
		b.shapes = append(b.shapes, shape)
		return nil
	}
	return errors.New("out of the shapesCapacity range")
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i < len(b.shapes) {
		return b.shapes[i], nil
	}
	return nil, errors.New("index doesn't exist or index went out of the range")

}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i < len(b.shapes) {
		s := b.shapes[i]
		b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
		return s, nil
	}
	return nil, errors.New("index doesn't exist or index went out of the range")
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i < len(b.shapes) {
		s := b.shapes[i]
		b.shapes[i] = shape
		return s, nil
	}
	return nil, errors.New("index doesn't exist or index went out of the range")
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	s := 0.0
	for i := 0; i < len(b.shapes); i++ {
		s += b.shapes[i].CalcPerimeter()
	}
	return s
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	s := 0.0
	for i := 0; i < len(b.shapes); i++ {
		s += b.shapes[i].CalcArea()
	}
	return s
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	c := false
	for i := 0; i < len(b.shapes); i++ {
		_, ok := b.shapes[i].(Circle)
		if ok {
			c = true
			if len(b.shapes) == 1 {
				b.shapes = b.shapes[:len(b.shapes)-1]
			}
			b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
			i = 0
		}
	}
	if c {
		return errors.New("circles are not exist in the list")
	}
	return nil
}
