// Initialize a MAP using a composite literal where the key is a string and the value is an int;
// print out the map; range over the map printing out just the key;
// range over the map printing out both the key and the value

package main

import "fmt"

func main() {
	m := map[string]int{
		"dog":  4,
		"cat":  4,
		"bird": 2,
	}

	fmt.Println(m)

	for i, _ := range m {
		fmt.Println(i)
	}

	fmt.Println()

	for i, v := range m {
		fmt.Println("Key:", i, "Value:", v)
	}
}
