package main

import "fmt"

func main() {
	s := circle{2}
	d := square{2, 5}
	shape := []Shapes{s, d}
	fmt.Println(shape[0].area())
	fmt.Println(shape[1].area())
}

type circle struct {
	radius int
}
type square struct {
	length  int
	breadth int
}

type Shapes interface {
	area() int
}

func (c circle) area() int {
	return c.radius * c.radius
}

func (s square) area() int {
	return s.length * s.breadth
}
