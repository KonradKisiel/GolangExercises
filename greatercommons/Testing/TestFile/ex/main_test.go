package main

import "testing"

func TestAvg(t *testing.T) {
	x := Avg(0, 10)
	if x != 5 {
		t.Error(x, "Is not average form 0 and 10")
	}
}
