package main

import (
	"io"
	"net/http"
)

func h(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "first link, provided by IHandler")
}

func h2(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "second link, provided by IHandler2")
}

func main() {
	// HandleFunc implements Handler interface
	http.HandleFunc("/godeeper/", h) // u can go deeper e.g first/something/someting / at the end of the line
	http.HandleFunc("/stayhere", h2) // u can't go deeper

	//http provide default Servmux, if handler is nil http use DefaultServMux
	http.ListenAndServe(":8080", nil)
}
