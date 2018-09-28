//putting value into URL
package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	//http://localhost:8080/?q=(searching word)
	v := req.FormValue("q")
	io.WriteString(w, "Do my search: "+v)
}
