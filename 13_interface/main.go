package main

import "math"
import "fmt"

type geometry interface {
	area() float64
	perm() float64
}

type Circle struct {
	x, y, r float64
}

type Rectange struct {
	w, h float64
}

func (r Rectange) area() float64 {
	return r.w * r.h
}

func (r Rectange) perm() float64 {
	return 2 * (r.w + r.h)
}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c Circle) perm() float64 {
	return 2 * math.Pi * c.r
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perm())
}

func main() {

	c := Circle{1, 1, 5}
	measure(c)

}
