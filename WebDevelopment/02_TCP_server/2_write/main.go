//Turn on telnet client in Windows:
//Go to Program and Functions -> Turn on and off Windows functions -> telnet CLient
//run this file
//in second terminal type: telnet localhost 8080

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
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()
	//we never get here
	//we have an open stream connection
	//how does the above rader know when it's done?
	fmt.Println("Code got here")
}
