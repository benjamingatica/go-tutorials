package main

import "testing"

func TestMySum(t *testing.T) {
	type Test struct {
		data     []int
		response int
	}

	tests := []Test{
		{[]int{1, 2, 3}, 6},
		{[]int{5, 2, 6}, 13},
		{[]int{6, 1, 6}, 13},
		{[]int{7, 8, 9}, 24},
	}

	for _, x := range tests {
		v := mySum(x.data...)
		resp := x.response
		if v != resp {
			t.Error("Expected", resp, "GOT", v)
		}
	}
}
