package main

import "testing"

func TestMySum(t *testing.T) {
	v := mySum(4, 6)
	if v != 10 {
		t.Error("Expected", 10, "GOT", v)
	}
}
