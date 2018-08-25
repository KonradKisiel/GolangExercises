package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	p := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(p), 12)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
	fmt.Println(string(bs))

	err = bcrypt.CompareHashAndPassword(bs, []byte(p))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Comparation validate")
	}
}
