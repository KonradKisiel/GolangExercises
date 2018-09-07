package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("---------------")
	//execute first file founded in container
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
	/* short version
	tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	fmt.Println("---------------")
	//execute first file founded in container
	tpl.Execute(os.Stdout, nil)
	*/
}
