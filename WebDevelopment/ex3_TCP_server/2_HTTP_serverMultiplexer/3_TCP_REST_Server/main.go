package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	// read request
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	fmt.Println()
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		//blink line means the end of a text
		if i == 0 {
			mux(conn, ln)
			// headers are done
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	// request line
	m := strings.Fields(ln)[0] // method
	u := strings.Fields(ln)[1] // uri
	fmt.Println("***METHOD***", m)
	fmt.Println("***URI***", u)
	// REST ROUTING
	if m == "GET" {
		switch u {
		case "/":
			index(conn)
		case "/about":
			about(conn)
		case "/contact":
			contact(conn)
		case "/apply":
			apply(conn)
		}
	}

	if m == "POST" {
		switch u {
		case "/apply":
			applyProcess(conn)
		}
	}
}

var links = `<a href="/">Index</a><br> 
<a href="/about">About</a><br>
<a href="/contact">Contact</a><br>
<a href="/apply">Apply</a><br>`

func template(conn net.Conn, innerBody string, title string) {

	wholeBody := `<!DOCTYPE html><html lang="en"><head><meta 
	charset="UTF-8"><title>` + title + `</title></head>
	<body>` + innerBody + `</body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Concept-Length: %d\r\n", len(wholeBody))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, wholeBody)
}

func index(conn net.Conn) {
	innerBody :=
		`<h1>Index</h1><br>` + links
	template(conn, innerBody, "Index")
}

func about(conn net.Conn) {
	innerBody :=
		`<h1>About</h1><br>` + links
	template(conn, innerBody, "About")
}

func contact(conn net.Conn) {
	innerBody :=
		`<h1>Contact</h1><br>` + links
	template(conn, innerBody, "Contact")
}

func apply(conn net.Conn) {
	innerBody :=
		`<h1>Apply</h1><br>` + links + `<form method="POST" action="/apply">
		<input type="submit" value="apply">
		</form>`
	template(conn, innerBody, "Apply")
}

func applyProcess(conn net.Conn) {
	innerBody :=
		`<h1>Applied Process</h1><br>` + links
	template(conn, innerBody, "Applied Process")
}
