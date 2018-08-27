package main

import "fmt"

func main() {
	c := make(chan int)
	//send
	go foo(c)

	//receive
	bar(c)

	fmt.Println("About to exit")
}

//send to
func foo(c chan<- int) {
	c <- 42
}

//receive from
func bar(c <-chan int) {
	fmt.Println(<-c)
}
