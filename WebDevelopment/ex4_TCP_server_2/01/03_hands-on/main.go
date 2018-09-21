/*
1. Take the previous program in the previous folder and change it so that:
* a template is parsed and served
* you pass data into the template
*/

package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func h(res http.ResponseWriter, req *http.Request) {
	tpl.Execute(res, "nothing to do here")
}

func h2(res http.ResponseWriter, req *http.Request) {
	tpl.Execute(res, "Woooffffff!!")
}

func h3(res http.ResponseWriter, req *http.Request) {
	tpl.Execute(res, "Elo")
}

func main() {
	http.HandleFunc("/", h)
	http.HandleFunc("/dog/", h2)
	http.HandleFunc("/me/", h3)

	http.ListenAndServe(":8080", nil)
}
