// Interface can be used to add behaviour of polymorphism
// Interface is collection of methods which type must have considering type implements interface.

package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	lenght, bredth float64
}

func (r rectangle) area() float64 {
	return r.bredth * r.lenght
}
func (r rectangle) perimeter() float64 {
	return 2 * (r.bredth + r.lenght)
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	rect := rectangle{lenght: 10, bredth: 50}
	cir := circle{radius: 30}

	var s shape = rect
	fmt.Printf("Area of rectangle is: %.2f\n", s.area())
	fmt.Printf("Perimeter of rectangle is: %.2f\n", s.perimeter())

	s = cir
	fmt.Printf("Area of circle is: %.2f\n", s.area())
	fmt.Printf("Perimeter of circle is: %.2f\n", s.perimeter())

}
