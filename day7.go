package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Triangle struct {
	firstLine  float64
	secondLine float64
	thirdLine  float64
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c Circle) Perimeter() float64 {
	return 2 * c.radius * math.Pi
}

func (r Rectangle) Area() float64 {
	return float64(r.height * r.width)
}

func (r Rectangle) Perimeter() float64 {
	return float64(2 * (r.height + r.width))
}

func (t Triangle) Perimeter() float64 {
	return t.firstLine + t.thirdLine + t.secondLine
}

func (t Triangle) Area() float64 {
	p := t.Perimeter() / 2
	area := math.Sqrt(p * (p - t.secondLine) * (p - t.firstLine) * (p - t.thirdLine))
	return area
}

func main() {
	c := Circle{radius: 120}
	r := Rectangle{height: 100, width: 100}
	t := Triangle{firstLine: 100, secondLine: 100, thirdLine: 100}
	fmt.Printf("Area of triangle %f.\n", t.Area())
	fmt.Printf("Perimeter of triangle %f.\n", t.Perimeter())
	fmt.Printf("Area of circle %f.\n", c.Area())
	fmt.Printf("Perimeter of circle %f.\n", c.Perimeter())
	fmt.Printf("Area of reactangle %f.\n", r.Area())
	fmt.Printf("Perimeter of reacangle %f.\n", r.Perimeter())
}
