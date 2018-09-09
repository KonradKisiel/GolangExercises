//Take the STRUCT “person” in the previous exercise and add a field called “favFood” that stores a slice of string;
//for the variable “p1” use a composite literal to add values to the field favFood; print out the values in favFood;
//also print out the values for “favFood” by using a for range loop

package main

import "fmt"

type person struct {
	fName   string
	lName   string
	favFood []string
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

	for _, v := range p1.favFood {
		fmt.Println(v)
	}
}
