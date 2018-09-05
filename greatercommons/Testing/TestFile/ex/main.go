package main

import "fmt"

func main() {
	fmt.Println(Avg(1, 10))
	fmt.Println(Avg(4, 6, 9))
}

func Avg(x ...float64) float64 {
	sum := 0.0
	for _, v := range x {
		sum += v
	}
	return sum / float64(len(x))
}
