//Adding onto this code: Can you assign “g1” to “x”? Why or why not?
package main

import "fmt"

type gator int

func main() {
	var g1 gator
	g1 = 42
	fmt.Println(g1)
	fmt.Printf("%T\n", g1)

	var x int
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	//example of conversion
	g1 = gator(x)
}
