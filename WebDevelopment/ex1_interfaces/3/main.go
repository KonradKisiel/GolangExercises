//create an interface type that both person and secretAgent implement
//declare a func with a parameter of the interface’s type
//call that func in main and pass in a value of type person
//call that func in main and pass in a value of type secretAgent

package main

import (
	"fmt"
)

type person struct {
	saying string
}

type secretAgent struct {
	person
	invisibility bool
}

func (p person) walking() {
	fmt.Println("Walking from point A to point B")
}

func (sa secretAgent) walking() {
	fmt.Println("Walking and trying to do not attend any observers, do not showing that I am a secret agent, changing directions and hiding if its possible, walking backwards in any kind of suspition of being watched")
}

type walk interface {
	walking()
}

func main() {
	p := person{
		saying: "Hello",
	}

	sa := secretAgent{
		person{
			"Hello, I am an ordinary person",
		},
		true,
	}

	fmt.Println()
	p.walking()
	fmt.Println()
	sa.walking()
	fmt.Println()
	sa.person.walking()
}
