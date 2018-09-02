package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("here is the error: %v", ce.info)
}

func main() {
	c1 := customErr{
		info: "custom error",
	}

	foo(c1)
}

func foo(e error) {
	//e.(customErr).info - assersion
	fmt.Println("foo ran -", e, "\n", e.(customErr).info)
	//fmt.Println("foo ran -", e, "\n")
}
