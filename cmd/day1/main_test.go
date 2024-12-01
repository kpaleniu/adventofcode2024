package main

import (
	"strings"
	"testing"
)

func TestDistance(t *testing.T) {
	test := `3   4
	4   3
	2   5
	1   3
	3   9
	3   3`

	a, b, err := parse(strings.NewReader(test))
	if err != nil {
		t.Error("failed to parse")
	}

	expected := 11
	distance := totalDistance(a, b)
	if distance != expected {
		t.Errorf("expected %d, got %d", expected, distance)
	}
}

func TestSimilarityScore(t *testing.T) {
	test := `3   4
	4   3
	2   5
	1   3
	3   9
	3   3`

	a, b, err := parse(strings.NewReader(test))
	if err != nil {
		t.Error("failed to parse")
	}

	expected := 31
	score := similarityScore(a, b)
	if score != expected {
		t.Errorf("expected %d, got %d", expected, score)
	}
}
