package Shapes

import (
	"math"
)

type Shape interface {
	CalculateArea() float64
	CalculatePerimeter() float64
}

type Circle struct {
	Radius float64
}

func (circle *Circle) CalculateArea() float64 {
	return 3.14 * circle.Radius * circle.Radius
}

func (circle *Circle) CalculatePerimeter() float64 {
	return 2 * 3.14 * circle.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (rectangle *Rectangle) CalculateArea() float64 {
	return rectangle.Width * rectangle.Height
}

func (rectangle *Rectangle) CalculatePerimeter() float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

type Triangle struct {
	FirstSide  float64
	SecondSide float64
	ThirdSide  float64
}

func (triangle *Triangle) CalculatePerimeter() float64 {
	return triangle.FirstSide + triangle.SecondSide + triangle.ThirdSide
}

func (triangle *Triangle) CalculateArea() float64 {
	var halfPerimeter = triangle.CalculatePerimeter() / 2
	return math.Sqrt(halfPerimeter * (halfPerimeter - triangle.FirstSide) * (halfPerimeter - triangle.SecondSide) * (halfPerimeter - triangle.ThirdSide))
}

type Square struct {
	Width float64
}

func (square *Square) CalculateArea() float64 {
	return square.Width * square.Width
}

func (square *Square) CalculatePerimeter() float64 {
	return 4 * square.Width
}
