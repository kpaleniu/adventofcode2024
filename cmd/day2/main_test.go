package main

import (
	"strings"
	"testing"
)

var input = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestSafety(t *testing.T) {
	count := part1(strings.NewReader(input))
	if count != 2 {
		t.Errorf("expected 2, got %d", count)
	}
}

func TestSafetyWithProblemDampener(t *testing.T) {
	count := part2(strings.NewReader(input))
	if count != 4 {
		t.Errorf("expected 4, got %d", count)
	}
}
func TestLevelsPart1(t *testing.T) {
	if !isSafe([]int{7, 6, 4, 2, 1}) {
		t.Error("all decreasing by 1 or 2")
	}
	if isSafe([]int{1, 2, 7, 8, 9}) {
		t.Error("large increase")
	}
	if isSafe([]int{9, 7, 6, 2, 1}) {
		t.Error("large decrease")
	}
	if isSafe([]int{1, 3, 2, 4, 5}) {
		t.Error("increasing and decreasing")
	}
	if isSafe([]int{8, 6, 4, 4, 1}) {
		t.Error("neither increase or decrease")
	}
	if !isSafe([]int{1, 3, 6, 7, 9}) {
		t.Error("all increasing by 1, 2 or 3")
	}
}
