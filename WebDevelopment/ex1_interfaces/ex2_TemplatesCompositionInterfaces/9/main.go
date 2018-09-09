//Give a method to both the “truck” and “sedan” types with the following signature
//transportationDevice() string
//Have each func return a string saying what they do.
//Create a value of type truck and populate the fields.
//Create a value of type sedan and populate the fields.
//Call the method for each value.

package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheels bool
}

type sedan struct {
	vehicle
	luxury bool
}

func (t truck) transportationDevice() string {
	return "I transport heavy things"
}

func (s sedan) transportationDevice() string {
	return "I a good transportation form for familly"
}

func main() {
	FordF150 := truck{
		vehicle{
			5,
			"blue",
		},
		true,
	}

	ToyotaCamry := sedan{
		vehicle{
			5,
			"silver",
		},
		false,
	}

	fmt.Println(ToyotaCamry.transportationDevice())
	fmt.Println()
	fmt.Println(FordF150.transportationDevice())
}
