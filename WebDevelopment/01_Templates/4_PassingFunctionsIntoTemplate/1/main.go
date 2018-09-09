package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

// create a FuncMap to register functions
// "uc" is what the func will be called in the template
// "uc" is the toUpper func from package strings
// "ft" is a func I declared
// "ft" slices a string, returning the first three characters

//type FuncMap map[string]interface{}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

var tpl *template.Template

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

type scholar struct {
	Name           string
	Accomplishment string
}

func init() {
	tpl = template.Must(template.New("tpl.gohtml").Funcs(fm).ParseFiles("tpl.gohtml"))
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
