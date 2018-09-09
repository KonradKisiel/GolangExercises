package main

import "fmt"

type gator int

type flamingo bool

type swampCreature interface {
	greeting()
}

func (f flamingo) greeting() {
	fmt.Println("Hello, I am pink and beautiful and wonderful.")
}

func (f gator) greeting() {
	fmt.Println("Hello, I am a gator")
}

func bayou(s swampCreature) {
	s.greeting()
}

func main() {
	var g1 gator
	var f1 flamingo
	g1.greeting()
	bayou(g1)
	bayou(f1)
}
