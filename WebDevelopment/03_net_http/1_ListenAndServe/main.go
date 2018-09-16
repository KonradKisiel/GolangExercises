package main

import (
	"fmt"
	"net/http"
)

type customType int

// implementing Handler interface for customType
func (m customType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want in this func")
}

func main() {
	var handler customType
	http.ListenAndServe(":8080", handler)
}
