package main

import (
	"fmt"
)

type customError struct {
	Error string
}

func main() {
	e := customError{
		Error: "error",
	}

	foo(e)
}

func foo(e customError) {
	fmt.Println(e.Error)
}
