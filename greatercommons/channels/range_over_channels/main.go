package main

import "fmt"

func main() {
	c := make(chan int)
	//send
	go foo(c)

	//receive
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("About to exit")
}

//send to
func foo(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

//receive from
func bar(c <-chan int) {
	fmt.Println(<-c)
}
