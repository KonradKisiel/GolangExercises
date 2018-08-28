package main

import (
	"fmt"
	"math/rand"
	"time"
)

//fan-in functionto let whoseover is ready talk
//Sending a channel on a channel, making gorutine waits its turn
//Receive all messages, then enable them again by sending on a private channel
//First we define a message type tht contains a channel for the reply
type Message struct {
	str  string
	wait chan bool
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	/*
		//Each speaker must wait for a go-ahead
		for i := 0; i < 5; i++ {
			msg1 := <-c
			fmt.Println(msg1.str)
			msg2 := <-c
			fmt.Println(msg2.str)
			msg1.wait <- true
			msg2.wait <- true
		}

		waitForIt := make(chan bool) //Shared between all messages

		c <- Message{fmt.Sprintf("%s: %d", msg, i), waitForIt}
		time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
		<-waitForIt
	*/
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving")
}

func boring(msg string) <-chan string { // Returns receive-only channel of strings.
	c := make(chan string)
	go func() { // We launch the gorutine from inside the function
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller
}
