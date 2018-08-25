package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	//fmt.Println(s)

	type person struct {
		First   string
		Last    string
		Age     int
		Sayings []string
	}

	var people []person
	err := json.Unmarshal([]byte(s), &people)
	if err != nil {
		fmt.Println(err)
	} else {
		//fmt.Println(people)
	}
	for _, v := range people {
		fmt.Println("Name:", v.First)
		fmt.Println("Surname:", v.Last)
		fmt.Println("Age:", v.Age)
		fmt.Println()
		fmt.Println("Says:")
		for _, x := range v.Sayings {
			fmt.Println("-", x)
		}
		fmt.Println("------------------------------------\n")
	}

}
