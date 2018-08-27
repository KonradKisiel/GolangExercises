package main

import "fmt"

func main() {
	//only sending values on to the channel
	c := make(chan<- int, 2)

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("---------------")
	fmt.Printf("%T\n", c)
}
