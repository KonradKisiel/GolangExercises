package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type scholar struct {
	Name           string
	Accomplishment string
}

type car struct {
	Manofacturer string
	Model        string
	Doors        int
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	Newton := scholar{
		Name:           "Izaak Newton",
		Accomplishment: "Foundations of physics, laws of motion",
	}

	Pitagoras := scholar{
		Name:           "Pitagoras",
		Accomplishment: "Fundations of geometry",
	}

	Einstein := scholar{
		Name:           "Einstein",
		Accomplishment: "Theory of relativity",
	}

	scholars := []scholar{
		Newton,
		Pitagoras,
		Einstein,
	}

	AMC := car{
		Manofacturer: "AMC",
		Model:        "AMX",
		Doors:        2,
	}

	AE86 := car{
		Manofacturer: "Toyota",
		Model:        "Corolla AE86",
		Doors:        2,
	}

	FridayZ := car{
		Manofacturer: "Datsun",
		Model:        "240z",
		Doors:        2,
	}

	cars := []car{
		AMC,
		AE86,
		FridayZ,
	}
	/*
		type items struct {
			Wisdom    []scholar
			Transport []car
		}

		data := items{
			Wisdom:    scholars,
			Transport: cars,
		}
	*/
	data := struct {
		Wisdom    []scholar
		Transport []car
	}{
		scholars,
		cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
