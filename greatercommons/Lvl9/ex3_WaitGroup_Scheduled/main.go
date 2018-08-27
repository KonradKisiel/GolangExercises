//race condition

/*
   Using goroutines, create an incrementer program
   have a variable to hold the incrementer value
   launch a bunch of goroutines
   each goroutine should
   read the incrementer value
   store it in a new variable
   yield the processor with runtime.Gosched()
   increment the new variable
   write the value in the new variable back to the incrementer variable
   use waitgroups to wait for all of your goroutines to finish
   the above will create a race condition.
   Prove that it is a race condition by using the -race flag
   if you need help, here is a hint: https://play.golang.org/p/FYGoflKQej
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup
	wg.Add(99)

	for i := 0; i < 99; i++ {
		go func() {
			runtime.Gosched()
			counter++
			fmt.Println(counter)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println()
	fmt.Println(counter)
}
