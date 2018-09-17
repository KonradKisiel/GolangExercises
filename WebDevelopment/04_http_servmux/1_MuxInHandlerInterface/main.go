package main

import (
	"io"
	"net/http"
)

type IHandler int

func (h IHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/about":
		io.WriteString(res, "about")
	case "/contact":
		io.WriteString(res, "contact")
	}
}

func main() {
	var h IHandler
	http.ListenAndServe(":8080", h)
}
