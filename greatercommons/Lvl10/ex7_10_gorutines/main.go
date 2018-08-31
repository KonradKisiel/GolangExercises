/*
   write a program that
   launches 10 goroutines
   each goroutine adds 10 numbers to a channel
   pull the numbers off the channel and print them
*/

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	c := make(chan int64)
	var mu sync.Mutex
	for i := 0; i < 10; i++ {
		go func() {
			mu.Lock()
			for j := 0; j < 10; j++ {
				atomic.AddInt64(&counter, 1)
				c <- counter
			}
			mu.Unlock()
		}()
	}
	/*
		do not work, because cannot assess channel range??
		for v := range c {
			fmt.Println(v)
		}*/
	for k := 0; k < 100; k++ {
		fmt.Println(<-c)
	}
	close(c)
	fmt.Println("Exit")
}
