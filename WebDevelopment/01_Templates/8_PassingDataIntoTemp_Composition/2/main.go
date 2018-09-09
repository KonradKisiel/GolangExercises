package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

type doubleOne struct {
	person
	Invisibility bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	p1 := doubleOne{
		person{
			Name: "Ian Fleming",
			Age: 56,
		},
		false,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", p1)
	if err != nil {
		log.Fatalln(err)
	}
}
