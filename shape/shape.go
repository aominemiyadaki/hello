package shape

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimetr() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimetr() float64 {
	return 2 * math.Pi * c.Radius
}

type Rect struct {
	W, H float64
}

func (r Rect) Area() float64 {
	return r.W * r.H
}

func (r Rect) Perimetr() float64 {
	return 2 * (r.W + r.H)
}

func PrintShape(s Shape) {
	// Добавили второй шаблон %.2f для периметра
	fmt.Printf("Площадь: %.2f, Периметр: %.2f\n", s.Area(), s.Perimetr())
}

type Triangle struct {
	A, B, C float64
}

func (t Triangle) Area() float64 {

	p := (t.A + t.B + t.C) / 2

	return math.Sqrt(p * (p - t.A) * (p - t.B) * (p - t.C))

}

func (t Triangle) Perimetr() float64 {
	return t.A + t.B + t.C
}
