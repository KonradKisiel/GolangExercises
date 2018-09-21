/*
ListenAndServe on port ":8080" using the default ServeMux.

Use HandleFunc to add the following routes to the default ServeMux:

"/"
"/dog/"
"/me/

Add a func for each of the routes.

Have the "/me/" route print out your name.
*/

package main

import (
	"io"
	"net/http"
)

func h(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "nothing to do here")
}

func h2(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Woooffffff!!")
}

func h3(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Elo")
}

func main() {
	http.HandleFunc("/", h)
	http.HandleFunc("/dog/", h2)
	http.HandleFunc("/me/", h3)

	http.ListenAndServe(":8080", nil)
}
