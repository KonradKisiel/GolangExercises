package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/satori/go.uuid"
)

/*
m := make(map[string]int)
m := map[string]int{}
someString, ok := m["someText"]
fmt.Println(someString, ok)

//if condition is ok print out someString
if someString, ok := m["someText"]; ok{
	fmt.Println(someString, ok)
}

if someString, ok := m["someText"]; !ok{
}
*/

type user struct {
	First    string
	Last     string
	Loggedin bool
}

var tpl *template.Template
var db = map[string]user{}

func init() {
	tpl = template.Must(template.ParseGlob("template/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID, err := uuid.NewV4()
		if err != nil {
			fmt.Println(err)
		}
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}
	var u user
	// process form submission
	if req.Method == http.MethodPost {
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		li := req.FormValue("loggedin") == "on"
		u = user{f, l, li}
		db[c.Value] = u
	}
	tpl.ExecuteTemplate(w, "index.gohtml", u)
	//io.WriteString(w, c.String())
}

func bar(w http.ResponseWriter, req *http.Request) {

	//get cookie
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		/*
			w.Header().Set("Location", "/")
			w.WriteHeader(http.StatusSeeOther)
		*/
		return
	}

	u, ok := db[c.Value]
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
	//io.WriteString(w, c.String())
}
