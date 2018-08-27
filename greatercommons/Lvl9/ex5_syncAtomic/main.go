/*
Fix the race condition you created in exercise #4 by using package atomic
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var counter int64
	wg.Add(99)
	for i := 0; i < 99; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println(atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println()
	fmt.Println(counter)
}
