package main

import (
	"fmt"
	"mystr"
	"strings"
)

const s = `“What the hell is that?" I laughed.
"It's my fox hat."
"Your fox hat?"
"Yeah, Pudge. My fox hat."
"Why are you wearing your fox hat?" I asked.
"Because no one can catch the motherfucking fox.” `

var xs []string

func main() {
	xs = strings.Split(s, " ")
	for _, v := range xs {
		fmt.Println(v)
	}

	fmt.Printf("\n%s\n", mystr.Cat(xs))
	fmt.Printf("\n%s\n\n", mystr.Join(xs))
}
