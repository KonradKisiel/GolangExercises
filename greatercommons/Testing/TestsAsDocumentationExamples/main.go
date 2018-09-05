// Package tstdoc
package tstdoc

// Sum adds an unlimmited number of values of type int
func Sum(xi ...int) int {
	s := 0
	for _, v := range xi {
		s += v
	}
	return s
}
