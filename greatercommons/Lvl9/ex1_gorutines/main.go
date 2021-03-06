/*
    in addition to the main goroutine, launch two additional goroutines
    each additional goroutine should print something out
    use waitgroups to make sure each goroutine finishes before your program exists
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)

	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	go func() {
		fmt.Println("Gorutine 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("Gorutine 2")
		wg.Done()
	}()

	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	wg.Wait()

	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	fmt.Println("Main thread")
}
