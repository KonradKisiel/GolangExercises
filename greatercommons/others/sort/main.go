package main

import (
	"fmt"
	"sort"
	//"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}
	Print(users)

	// your code goes here

	sort.Sort(byAge(users))
	Print(users)

	sort.Sort(byLast(users))
	Print(users)

	//sorting by sayings
	for _, u := range users {
		sort.Strings(u.Sayings)
	}
	Print(users)
}

func Print(u []user) {
	fmt.Println("***************************************************")
	for _, v := range u {
		fmt.Println("---------------------------------------------------")
		fmt.Println("Name:", v.First, v.Last)
		fmt.Println("Age:", v.Age)
		fmt.Println("Sayings:")
		for _, x := range v.Sayings {
			fmt.Printf("\t%v\n", x)
		}
		fmt.Println("---------------------------------------------------")
	}
	fmt.Println("***************************************************")
	fmt.Println()
}

type byAge []user

func (a byAge) Len() int           { return len(a) }
func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a byAge) Swap(i, j int)      { a[i].Age, a[j].Age = a[j].Age, a[i].Age }

type byLast []user

func (l byLast) Len() int           { return len(l) }
func (l byLast) Less(i, j int) bool { return l[i].Last < l[j].Last }
func (l byLast) Swap(i, j int)      { l[i].Last, l[j].Last = l[j].Last, l[i].Last }
