package main

import (
	"net/http"
)

func main() {
	// index.html is a special case, index instead of files in darectory will be not shown
	http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}
