//Add a method to type “person” with the identifier “walk”.
//Have the func return this string: <person’s first name> is walking.
//Remember, you make a func a method by giving the func a receiver.

package main

import "fmt"

type person struct {
	fName   string
	lName   string
	favFood []string
}

func (p person) walk() string {
	return fmt.Sprintln(p.fName, "is walking")
}

func main() {
	p1 := person{
		fName: "Alan",
		lName: "Turing",
		favFood: []string{
			"apples",
			"beef",
			"pancakes",
			"spaghetti",
		},
	}

	s := p1.walk()
	fmt.Println(s)
}
