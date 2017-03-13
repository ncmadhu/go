package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2 * r.width + 2 * r.height
}

func (c circle) area() float64 {
	return 2 * math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Printf("area:%f\n", g.area())
	fmt.Printf("perimeter:%f\n", g.perimeter())
}

func main() {
	r := rectangle{2,3}
	c := circle{5}
	measure(r)
	measure(c)
}
