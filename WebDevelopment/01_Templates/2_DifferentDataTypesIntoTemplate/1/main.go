package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	template1, err := template.ParseFiles("template1.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = template1.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
