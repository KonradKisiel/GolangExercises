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

type user struct {
	Name  string
	Motto string
	Admin bool
}

func main() {

	u1 := user{
		Name:  "John",
		Motto: "But there is only one thing which gathers people into seditious commotion, and that is oppression.",
		Admin: true,
	}

	u2 := user{
		Name:  "Libertine",
		Motto: "Let's set the world on fire",
		Admin: false,
	}

	u3 := user{
		Name:  "",
		Motto: "Anybody",
		Admin: true,
	}

	users := []user{u1, u2, u3}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", users)
	if err != nil {
		log.Fatalln(err)
	}
}
