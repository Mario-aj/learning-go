package main

import "math"

func Perimeter(rectangle Rectangle) float64 {

	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Height * rectangle.Width
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Ray float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Ray * c.Ray
}

type Form interface {
	Area() float64
}
