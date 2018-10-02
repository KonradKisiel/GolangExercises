// Using cookies, track how many times a user has been to your website domain
package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", createCookies)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

var VisitsCounter = 0

func createCookies(w http.ResponseWriter, req *http.Request) {
	VisitsCounter++
	http.SetCookie(w, &http.Cookie{
		Name:  "cookie",
		Value: "visit nr: " + strconv.Itoa(VisitsCounter),
	})
	fmt.Println("visit nr: " + strconv.Itoa(VisitsCounter))
}
