/*
Fix the race condition you created in the previous exercise by using a mutex
    it makes sense to remove runtime.Gosched()
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup
	var m sync.Mutex
	wg.Add(99)

	for i := 0; i < 99; i++ {
		go func() {
			m.Lock()
			//	runtime.Gosched()
			counter++
			fmt.Println(counter)
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println()
	fmt.Println(counter)
}
