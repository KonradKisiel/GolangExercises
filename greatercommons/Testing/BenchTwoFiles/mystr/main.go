package mystr

import "strings"

func Cat(xs []string) string {
	s := xs[0]
	/*
		for i := 1; i < len(xs); i++ {
			s += " "
			s += xs[i]
		}*/
	for _, v := range xs[1:] {
		s += " "
		s += v
	}
	return s
}

func Join(xs []string) string {
	return strings.Join(xs, " ")
}
