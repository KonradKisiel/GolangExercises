package main

import (
	"fmt"
	"log"
)

type norgateMathError struct {
	lat  string
	long string
	err  error
}

func (n norgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occured: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		//creating errors.New with fotmat printing, for more detaied informations
		nme := fmt.Errorf("norgate math redux: square root of negative number: %v", f)
		return 0, norgateMathError{"50.2289N", "99.4656 W", nme}
	}
	return f * f, nil
}

//use of structs with error type in standars library
//
//golang.org/pkg/net/#OpError
//golang.org/src/pkg/net/dial.go
//golang.org/src/pkg/net/net.go
//
//golang.org/src/pkg/encoding/json/decode.go
