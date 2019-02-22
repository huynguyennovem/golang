package main

import (
	"fmt"
	"math"
)

// a "receiver" - Circle is present to represent the container of the function - area().
// This receiver can be used to call a function using "." operator.

type Circle struct {
	x, y, radius float64
}

func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func main() {
	circle := Circle{x: 1.2, y: 1.3, radius: 3}
	fmt.Printf("%f", circle.area())
}
