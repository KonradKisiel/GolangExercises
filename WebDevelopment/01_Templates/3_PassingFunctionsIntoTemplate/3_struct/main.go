package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type scholar struct{
	Name string
	Accomplishment string

}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

    Newton := scholar{
		Name: "Izaak Newton",
		Accomplishment: "Foundations of physics, laws of motion",
	}

/*
	scholars := map[string]string{
		"Greece":   "Aristotele",
		"GB":       "Newton",
		"American": "Einstein",
		"Chroatia": "Tesla"}
*/
		
	err := tpl.Execute(os.Stdout, scholars)
	if err != nil {
		log.Fatalln(err)
	}
}
