package main

import (
	"fmt"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("err happened", err)
		//log.Println("err happened", err) //log added additional time stamp
		//log.Fatalln(err) //time stamp + force exit -> os.Exit(1), closes all func
		panic(err) //stops execution of current gorutine, other function run
	}
}

func foo() {
	fmt.Println("Deffer func")
}
