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

	err := tpl.Execute(os.Stdout, scholars)
	if err != nil {
		log.Fatalln(err)
	}
}
