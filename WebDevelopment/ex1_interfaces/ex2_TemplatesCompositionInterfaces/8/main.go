//Using the vehicle, truck, and sedan structs: using a composite literal, create a value of type truck and assign values to the fields;
//using a composite literal, create a value of type sedan and assign values to the fields. Print out each of these values.
//Print out a single field from each of these values.

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

	fmt.Println(FordF150)

	fmt.Println(ToyotaCamry)

	fmt.Println(FordF150.color)

	fmt.Println(ToyotaCamry.luxury)
}
