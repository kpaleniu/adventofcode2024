package main

import (
	"strings"
	"testing"
)

var s = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func TestDescription(t *testing.T) {
	c := part1(s)
	if c != 18 {
		t.Errorf("wrong count, got %d", c)
	}
}

func TestRotate(t *testing.T) {
	foo := `XS
MA
AM
SX`
	l := strings.Split(foo, "\n")
	r := rotate(runeSlices(l))
	xmas := string(r[0])
	if xmas != "XMAS" {
		t.Errorf("wrong, got %s", xmas)
	}
	samx := string(r[1])
	if samx != "SAMX" {
		t.Errorf("wrong, got %s", samx)
	}
}

func TestRotate2(t *testing.T) {
	foo := `ABC
DEF
GHI`
	l := strings.Split(foo, "\n")
	r := rotate(runeSlices(l))
	u := string(r[0])
	if u != "GDA" {
		t.Errorf("wrong, got %s", u)
	}
	u = string(r[1])
	if u != "HEB" {
		t.Errorf("wrong, got %s", u)
	}
	u = string(r[2])
	if u != "IFC" {
		t.Errorf("wrong, got %s", u)
	}
}
func TestDiagonal(t *testing.T) {
	foo := `ABC
DEF
GHI`
	l := strings.Split(foo, "\n")
	r := runeSlices(l)

	flattened := make([]rune, 0)
	for r := range diagonal(r) {
		flattened = append(flattened, r)
	}

	result := string(flattened)
	expected := "A DB GEC HF I"
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestPart2(t *testing.T) {
	foo := `M.S
.A.
M.S`
	result := part2(foo)
	if result != 1 {
		t.Errorf("expected 1, got %d", result)
	}

	bar := `.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........`

	result = part2(bar)
	if result != 9 {
		t.Errorf("expected 9, got %d", result)
	}
}
