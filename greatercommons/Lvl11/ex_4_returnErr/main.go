package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	se := sqrtError{
		lat:  "50.2289 N",
		long: "99.4656 W",
		err:  fmt.Errorf("Cannot squre numbers less than 0"),
	}
	_, err := sqrt(-10.23, se)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64, se sqrtError) (float64, error) {
	if f < 0 {
		// write your error code here
		return 0, se.err
	}
	return f * f, nil
}
