// Initialize a SLICE of int using a composite literal;
// print out the slice;
// range over the slice printing out just the index;
// range over the slice printing out both the index and the value

package main

import "fmt"

func main() {
	s := []int{1, 4, 6, 1, 7}

	fmt.Println(s)

	for index, _ := range s {
		fmt.Println(index)
	}

	for i, v := range s {
		fmt.Println("index:", i, "value", v)
	}
}
