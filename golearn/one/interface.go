package main

import (
	"fmt"
	"math"
)

//interface are named collections of method signatures

type geometry interface {
	area() float64
	perim() float64
}

type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}
func (r rectangle) perim() float64 {
	return 2*r.height + 2*r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func Interface() {
	r := rectangle{4, 5}
	c := circle{6}

	measure(r)
	measure(c)
}
