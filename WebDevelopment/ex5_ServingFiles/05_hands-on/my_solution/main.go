/*
# Serve the files in the "starting-files" folder

## To get your images to serve, use only this:

``` Go
	fs := http.FileServer(http.Dir("public"))
```

Hint: look to see what type FileServer returns, then think it through.
*/
package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("starting-files/templates/index.gohtml"))
}

func main() {
	fs := http.FileServer(http.Dir("public"))

	http.Handle("/resources", fs)
	http.HandleFunc("/", start)

	http.ListenAndServe(":8080", nil)
}

func start(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
}
