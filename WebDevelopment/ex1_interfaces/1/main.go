// HANDS ON 1
// create a type square
// create a type circle
// attach a method to each that calculates area and returns it
// create a type shape which defines an interface as anything which has the area method
// create a func info which takes type shape and then prints the area
// create a value of type square
// create a value of type circle
// use func info to print the area of square
// use func info to print the area of circle

package main

import "fmt"

type square struct {
	s float64
}

type circle struct {
	r float64
}

func (sq square) area() float64 {
	return sq.s * sq.s
}

func (c circle) area() float64 {
	return 3.14 * c.r * c.r
}

type shape interface {
	area() float64
}

func info(sh shape) {
	fmt.Println(sh.area())
}

func main() {
	sq := square{
		s: 5.5,
	}

	c := circle{
		r: 6.2,
	}

	info(sq)
	info(c)
}
