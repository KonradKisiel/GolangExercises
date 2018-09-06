// create a struct that holds person fields
// create a struct that holds secret agent fields and embeds person type
// attach a method to person: pSpeak
// attach a method to secret agent: saSpeak
// create a variable of type person
// create a variable of type secret agent
// print a field from person
// run pSpeak attached to the variable of type person
// print a field from secret agent
// run saSpeak attached to the variable of type secret agent
// run pSpeak attached to the variable of type secret agent

package main

import "fmt"

type person struct {
	saying string
}

type secretAgent struct {
	person
	invisibility bool
}

func (p person) pSpeak() {
	fmt.Println(p.saying)
}

func (sa secretAgent) saSpeak() {
	fmt.Println(sa.person.saying)
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

	fmt.Println(p.saying)
	p.pSpeak()

	fmt.Println(sa.person.saying)
	sa.saSpeak()
}
