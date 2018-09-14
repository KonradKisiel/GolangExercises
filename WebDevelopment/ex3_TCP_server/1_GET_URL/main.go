// create a server that returns the URL of the GET request
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
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

	// write response
	respond(conn)
}

func request(conn net.Conn) {
	var str2 string
	i := 0
	scanner := bufio.NewScanner(conn)
	fmt.Println()
	for scanner.Scan() {
		ln := scanner.Text()
		str1 := ln[6:]

		//blink line means the end of a text
		if i > 0 {
			fmt.Println(str1 + str2)
			// headers are done
			break
		}
		i++
		str2 = ln[4 : len(ln)-8]
	}
}

// HTTP response
func respond(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta 
	charset="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Concept-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
