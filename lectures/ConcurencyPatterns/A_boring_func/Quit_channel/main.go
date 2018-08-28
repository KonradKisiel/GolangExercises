package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	quit := make(chan bool)
	c := boring("Joe", quit)
	timeout := time.After(5 * time.Second)
	for i := rand.Int(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- true
}

func boring(msg string, quit chan bool) <-chan string { // Returns receive-only channel of strings.
	c := make(chan string)
	select {
	case c <- fmt.Sprintf("%s: %d", msg, i):
	case <-quit:
		return
	}
}
