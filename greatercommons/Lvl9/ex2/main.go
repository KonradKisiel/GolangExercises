package main

import "fmt"

type person struct {
	say string
}

func (p *person) speak() {
	fmt.Println(p.say)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		say: "Hello, I'm p1",
	}
	saySomething(&p1)

	// does not work
	// saySomething(p1)

	p1.speak()
}
