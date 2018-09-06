package main

import "fmt"

//custom type
type hotdog int

// variaable declaration out of scope
var name string = "John"

//composite literal struct
type person struct {
	//field not visible outside of package
	fname string
	//field visible outside of package
	Lname string
}

//composition
type secretAgent struct {
	person
	licenceToKill bool
}

func (p person) speak() {
	fmt.Println(p.fname, `says, "Good morning, James."`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.Lname, `says, "Shaken, not stirred."`)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	// short operand inner scope declaration of float64
	i := 6.0

	//composite literals
	// slice of ints
	xi := []int{2, 4, 7, 9, 42}
	// map
	m := map[string]int{
		"Todd": 45,
		"Job":  42,
	}

	fmt.Printf("%T:\t%v\n%T:\t%v\n%T:\t%v\n", i, i, xi, xi, m, m)

	p1 := person{
		"Miss",
		"Moneypenny",
	}

	p1.speak()
	fmt.Println(p1)

	sa1 := secretAgent{
		person{
			"James",
			"Bond"},
		true,
	}
	sa1.speak()

	//polimorphism, using interface for types implementing speak()
	saySomething(p1)
	saySomething(sa1)
}
