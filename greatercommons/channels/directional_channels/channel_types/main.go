package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(<-chan int) //receive
	cs := make(chan<- int) // send

	fmt.Printf("c\t%T\n", c)
	fmt.Printf("c\t%T\n", cr)
	fmt.Printf("c\t%T\n", cs)

	/*
		//specyfic to general doesn't assign
		c = cr
		c = cs

		//specyfic to specific doesn't assign
		cs = cr
	*/

	//general to specific assigns
	cs = c
	cs = c

	//general to specific converts
	fmt.Printf("c\t%T\n", (<-chan int)(c))
	fmt.Printf("c\t%T\n", (chan<- int)(c))
}
