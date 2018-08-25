package main

import (
	"fmt"
	"io"
	"os"
)

//declaring a struct for handling received data
type fileWriter struct {
}

//implementing Write interface for print the file content to the terminal
func (fileWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}

func main() {
	// selecting a file name form console input
	name := os.Args[1]
	// opening a file
	file, err := os.Open(name)
	// checking for errors
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// struct initalization
	fw := fileWriter{}
	// copying data from file to struct
	io.Copy(fw, file)
}
