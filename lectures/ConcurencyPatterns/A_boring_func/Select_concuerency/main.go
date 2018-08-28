package main

import (
	"fmt"
)

/*The select statement provides another way to handle multiple channels.
It's like a switch, but each case is a communication:
	-All channels are evaluated
	-Selection blocks until one communication can proceed, which then does
	-If multiple cam proceed, select chooses pseudo-randomly
	-A default clase, if present, execiutes immediately if no channel is ready
*/
func main() {
	select {
	case v1 := <-c1:
		fmt.Printf("received %v from c1\n", v1)
	case v2 := <-c2:
		fmt.Printf("received %v from c2\n", v1)
	case c3 <- 23:
		fmt.Printf("send %v to c3\n", 23)
	default:
		fmt.Printf("no one was ready to communicate\n")
	}
}
