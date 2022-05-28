package golang_united_school_homework

import (
	"errors"
)

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

	return errors.New("box capacity exceeded")
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i < len(b.shapes) {
		return b.shapes[i], nil
	}

	return nil, errors.New("shape no such")
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	shape, err := b.GetByIndex(i)
	if err == nil {
		b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
		return shape, nil
	}

	return nil, err
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	shapeOld, err := b.GetByIndex(i)
	if err == nil {
		b.shapes[i] = shape
		return shapeOld, nil
	}

	return nil, err
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	perimeter := 0.0
	for _, shape := range b.shapes {
		perimeter += shape.CalcPerimeter()
	}

	return perimeter
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	area := 0.0
	for _, shape := range b.shapes {
		area += shape.CalcArea()
	}

	return area
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	shapes := make([]Shape, 0, len(b.shapes))
	for _, shape := range b.shapes {
		if _, ok := shape.(*Circle); !ok {
			shapes = append(shapes, shape)
		}
	}
	
	if len(b.shapes) == len(shapes) {
		return errors.New("circles are not exist in the list")
	}

	b.shapes = shapes

	return nil
}
