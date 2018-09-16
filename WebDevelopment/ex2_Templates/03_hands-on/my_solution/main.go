package main

import (
	"log"
	"os"
	"text/template"
)

type Hotel struct {
	Name   string
	Adress string
	City   string
	Zip    string
	Region string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	Hotels := []Hotel{
		Hotel{
			Name:   "California",
			Adress: "5th Avenue",
			City:   "Los Angeles",
			Zip:    "34-560",
			Region: "Southern",
		},
		Hotel{
			Name:   "Name1",
			Adress: "Adress1",
			City:   "City1",
			Zip:    "Zip1",
			Region: "Central",
		},
	}

	err := tpl.Execute(os.Stdout, Hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
