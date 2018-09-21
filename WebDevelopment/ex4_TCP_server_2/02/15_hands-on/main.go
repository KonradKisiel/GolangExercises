/*
# Building upon the code from the previous problem:

Change your RESPONSE HEADER "content-type" from "text/plain" to "text/html"

Change the RESPONSE from "CHECK OUT THE RESPONSE BODY PAYLOAD" (and everything else it contained: request method, request URI) to an HTML PAGE that prints "HOLY COW THIS IS LOW LEVEL" in <h1> tags.
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		serve(conn)
	}
}

func serve(c net.Conn) {
	defer c.Close()
	scanner := bufio.NewScanner(c)
	var ln string
	i := 0
	for scanner.Scan() {
		ln = scanner.Text()
		fmt.Println(ln)

		//blink line means the end of a text
		if i == 0 {
			// headers are done
			break
		}
		i++

	}

	m := strings.Fields(ln)[0] // method
	u := strings.Fields(ln)[1] // uri
	fmt.Println("*Method:", m)
	fmt.Println("****URI:", u)
	fmt.Println()

	body := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Index</title>
	</head>
	<body>
		<h1>HOLY COW THIS IS LOW LEVEL</h1>

		<p>Method:`+m+`</p>
		<p>URI: `+u+`</p>
	</body>
	</html>`

	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)

}
