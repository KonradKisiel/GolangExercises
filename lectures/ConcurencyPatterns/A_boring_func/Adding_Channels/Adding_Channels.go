package main

import (
	"fmt"
	"math/rand"
	"time"
)

// channel operations are blocking, to archive synchronization of receiver and sender (<-c and c<-)

func main() {
	c := make(chan string)
	go boring("boring!", c)

	for i := 0; i < 5; i++ {
		// <-c means wainting for value c to execute the rest of the code
		fmt.Printf("You say: %q\n", <-c) // Receive expression is just a value
	}

	fmt.Println("You're boring; I'm leaving")
}

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		// c<- means that the functon waits to receive c value
		c <- fmt.Sprintf("%s %d", msg, i) // Expression to be sent can be ny suitable value
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
