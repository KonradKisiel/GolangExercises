package main

import (
	"dog" //package included in src directory (GOROOT or GOPATH/src)
	"fmt"
)

/*
//Packge dog
package dog

//Yers take as an input nuber of human yers and return dog years
func Years(humanYears int) int {
	return 7 * humanYears
}

*/

func main() {
	fmt.Println(dog.Years(10))
}
