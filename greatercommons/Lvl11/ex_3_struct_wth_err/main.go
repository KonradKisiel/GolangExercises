package main

import (
	"errors"
	"fmt"
)

type customErr struct {
	err error
}

func main() {
	newErr := errors.New("Elo form custom error")
	custom := customErr{
		err: newErr,
	}
	foo(custom)
}

func foo(c customErr) error {
	fmt.Println(c.err)
	return c.err
}
