//Create a program that prints out your OS and ARCH. Use the following commands to run it

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("CPU:", runtime.GOARCH, "\nOS:", runtime.GOOS)
}
