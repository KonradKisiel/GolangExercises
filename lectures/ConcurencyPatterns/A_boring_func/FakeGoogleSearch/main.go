package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Search func(query string) Result

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

//to avoid timeot, we replicate the servers
//send request to multiple replicas, and use the first response

func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

func Google(query string) (results []Result) {
	c := make(chan Result)
	//fanIn design pattern (no locks, no callbacks)
	//reduce tail latency using replicated search servers
	go func() { c <- Web(query, Web1, Web2) }()
	go func() { c <- Image(query, Image1, Image2) }()
	go func() { c <- Video(query, Video1, Video2) }()

	//dont wait for slow servers
	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("time out")
			return
		}
	}
	return
}

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := First("golang",
		fakeSearch("replica1"),
		fakeSearch("replica2"))
	elapsed := time.Since(start)
	fmt.Println(result)
	fmt.Println(elapsed)
}
