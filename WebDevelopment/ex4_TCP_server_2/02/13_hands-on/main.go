/*
# Building upon the code from the previous problem:

Print to standard out (the terminal) the REQUEST method and the REQUEST URI from the REQUEST LINE.

Add this data to your REPONSE so that this data is displayed in the browser.
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

	body := "CHECK OUT THE RESPONSE BODY PAYLOAD\n\nMethod:" + m + "\nURI: " + u

	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/plain\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)

}
