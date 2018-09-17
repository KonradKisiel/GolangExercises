package main

import (
	"io"
	"net/http"
)

type IHandler int

func (h IHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "first link, provided by IHandler")
}

type IHandler2 int

func (h2 IHandler2) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "second link, provided by IHandler2")
}

func main() {
	var h IHandler
	var h2 IHandler2

	mux := http.NewServeMux()
	mux.Handle("/godeeper/", h) // u can go deeper e.g first/something/someting / at the end of the line
	mux.Handle("/stayhere", h2) // u can't go deeper

	http.ListenAndServe(":8080", mux)
}
