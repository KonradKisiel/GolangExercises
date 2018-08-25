// race condition gorutines
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Arch:", runtime.GOARCH)
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Gorutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100 // nr of gorutines
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			//if code is locked, noone other gorutine van access that code
			mu.Lock()
			v := counter
			runtime.Gosched() //optional??
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		//fmt.Println("Gorutines:", runtime.NumGoroutine())
		//fmt.Println("Count:", counter)
	}
	wg.Wait()
	fmt.Println("Gorutines:", runtime.NumGoroutine())
	fmt.Println("Count:", counter)
}
