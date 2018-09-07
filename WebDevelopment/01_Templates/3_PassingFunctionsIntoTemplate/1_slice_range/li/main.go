package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	scholars := []string{"Aristotele", "Pitagoras", "Platon", "Newton", "Einstein", "Tesla"}
	err := tpl.Execute(os.Stdout, scholars)
	if err != nil {
		log.Fatalln(err)
	}
}
