package main

import (
	"fmt"
	"net/http"
)

type IHandler int

// Handler interface implementation
func (h IHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Entry-Key", "this is form the first line of header")
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(res, "<h1>Any code you want in this func</h1>")
}

func main() {
	var h IHandler
	http.ListenAndServe(":8080", h)
}
