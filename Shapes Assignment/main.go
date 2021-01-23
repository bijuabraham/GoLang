package main

import (
	"fmt"
)

type Areas interface {
	area() float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func main() {
	t := triangle{height: 2.5, base: 3.5}
	printArea(t)

	s := square{sideLength: 2.5}
	printArea(s)
}

func (s square) area() float64 {
	return s.sideLength * s.sideLength
}

func (r triangle) area() float64 {
	return r.base * r.height
}

func printArea(a Areas) {
	fmt.Println(a)
	fmt.Printf("area     : %v \n", a.area())
}
