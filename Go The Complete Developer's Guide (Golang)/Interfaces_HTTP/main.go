package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

//now logWriter implementing Write interface
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Lenght of slice:", len(bs), "bytes")
	return len(bs), nil
}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		//log.Fatal(err)
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	//make slice with 9999 capacity
	/*
		bs := make([]byte, 9999)
		resp.Body.Read(bs)
		fmt.Println(string(bs))
	*/

	lw := logWriter{}
	io.Copy(lw, resp.Body)

	//Stdout is a value which is exported from os package, IT is type of File
	//File is a type which is implements Read interface
	//io.Copy(os.Stdout, resp.Body)
}
