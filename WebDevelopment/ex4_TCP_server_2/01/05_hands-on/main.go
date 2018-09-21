/*
1. Take the previous program and change it so that:
* func main uses http.Handle instead of http.HandleFunc
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

type IHandler int

func (h IHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	tpl.Execute(res, "nothing to do here")
}

type IHandler2 int

func (h2 IHandler2) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	tpl.Execute(res, "Woooffffff!!")
}

type IHandler3 int

func (h3 IHandler3) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	tpl.Execute(res, "Elo")
}

func main() {
	var h IHandler
	var h2 IHandler2
	var h3 IHandler3

	mux := http.NewServeMux()
	mux.Handle("/", h)
	mux.Handle("/dog/", h2)
	mux.Handle("/me/", h3)

	http.ListenAndServe(":8080", mux)
}
