package main

import (
	"fmt"
	"math"
	_ "math"
)

type square struct{
	l float64
}

type circle struct {
	r float64
}

func (c circle) area() float64 {
	return c.r *c.r * math.Pi
}
func (s square) area() float64{
	return s.l * s.l
}

type shape interface {
	area() float64
}

type shapeFunc func() float64

func (s shapeFunc) area() float64 {
	return s()
}

func info(s shape) {
	fmt.Printf("the area is: %v\n", s.area())
}

func main() {
	c := circle{
		r:3,
	}
	fmt.Println(c)
	info(c)

	s := square{
		l: 5,
	}
	fmt.Println(s)
	info(s)

	info(shapeFunc(func() float64{
		return 10
	}))
}
