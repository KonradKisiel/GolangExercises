package main

import "testing"

func TestMySum(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{21, 21}, 42},
		test{[]int{3, 4}, 7},
		test{[]int{12, 32}, 44},
		test{[]int{-51, 10}, -41},
	}

	//testing function mySum() using tests table
	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("Expected:", v.answer, ", got:", x)
		}
	}
}
