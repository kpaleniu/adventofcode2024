package main

import "testing"

func TestDescription(t *testing.T) {
	s := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	if part1(s) != 161 {
		t.Error("incorrect result")
	}
}
