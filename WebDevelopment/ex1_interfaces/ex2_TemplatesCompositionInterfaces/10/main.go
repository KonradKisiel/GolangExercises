//Create a new type called “transportation”.
//The underlying type of this new type is interface.
//An interface defines functionality.
//Said another way, an interface defines behavior.
//For this interface, any other type that has a method with this signature …
//transportationDevice() string
//… will automatically, implicitly implement this interface.
//Create a func called “report” that takes a value of type “transportation” as an argument.
//The func should print the string returned by “transportationDevice()” being called for any type that implements the “transportation” interface.
//Call “report” passing in a value of type truck. Call “report” passing in a value of type sedan.

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

type transportation interface {
	transportationDevice() string
}

func report(t transportation) {
	fmt.Println(t.transportationDevice())
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

	report(FordF150)
	fmt.Println()
	report(ToyotaCamry)
}
