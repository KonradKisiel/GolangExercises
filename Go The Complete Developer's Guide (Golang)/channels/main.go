package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.stackoverflow.com",
		"http://www.golang.org",
		"http://www.amazon.com",
	}

	//channels are using to communicate throught different gorutines
	//creating channel
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
		//go fmt.Println(<-c)
	}
	/*
			//printout channel data
			fmt.Println(<-c)
			fmt.Println(<-c)
			fmt.Println(<-c)
			fmt.Println(<-c)
			fmt.Println(<-c)

		for {
			go checkLink(<-c, c)
		}
	*/

	for l := range c {
		//go checkLink(l, c)
		//function literal(in other languages lambda/anonymous function)

		// NEVER TRY TO ACCES THE SAME VARIABLE WITH DIFFERENT CHILD RUTINE
		// to avoid this, we passing value as function literal argument
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) //() is necesary after function literal, passing l as a function argument
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + "might be down!")
		//sending a messagge into a channel
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
