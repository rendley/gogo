package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (r Rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Printf("%T: %+v\n", g, g)
	fmt.Println("area:", g.area())
	fmt.Println("perimetr:", g.perim())
}

func main() {
	r := Rect{10, 5}
	c := Circle{radius: 14}
	measure(r)
	fmt.Println("##############")
	measure(c)
}
