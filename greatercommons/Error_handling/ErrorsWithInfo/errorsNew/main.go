package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrNorgateMath = errors.New("norgate math: square root of negative number")

func main() {
	fmt.Printf("%T\n", ErrNorgateMath)
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		//	return 0, errors.New("norgate math: square root of negative number")
		return 0, ErrNorgateMath
	}
	return f * f, nil
}

//use of errors.New in standard library
//golang.org/src/pkg/bufio/bufio.go
//golang.org/src/pkg/io/io.go
