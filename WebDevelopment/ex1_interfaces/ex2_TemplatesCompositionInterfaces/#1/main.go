/*
#1
CREATE TYPE PERSON
THE UNDERLYING TYPE IS A STRUCT
FIELDS: FIRST, LAST, AGE

#2
USE A COMPOSITE LITERAL
TO CREATE A VALUE OF TYPE PERSON
AND ASSIGN IT TO A VARIABLE
USING THE SHORT DECLARATION OPERATOR
*/

package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		"John",
		"Locke",
		21,
	}

	fmt.Println(p1)
}
