package main

import (
	"fmt"
	"os"
	"strings"
)

func runeSlices(slice []string) [][]rune {
	runes := make([][]rune, len(slice))
	for i := range slice {
		runes[i] = make([]rune, len(slice[i]))
		for j, c := range slice[i] {
			runes[i][j] = c
		}
	}
	return runes
}

func rotate(slice [][]rune) [][]rune {
	xl, yl := len(slice[0]), len(slice)
	r := make([][]rune, xl)
	for i := range r {
		r[i] = make([]rune, yl)
	}
	// 1: transpose
	for y := range yl {
		for x := range xl {
			r[x][y] = slice[y][x]
		}
	}
	// 2: reverse rows
	for y := range xl {
		for x := range yl / 2 {
			r[y][x], r[y][yl-x-1] = r[y][yl-x-1], r[y][x]
		}
	}

	return r
}

func diagonal(runes [][]rune) func(func(rune) bool) {
	xl, yl := len(runes[0]), len(runes)
	return func(yield func(rune) bool) {
		for y := 0; y < yl; y++ {
			for x := y; x >= 0; x-- {
				if !yield(runes[x][y-x]) {
					return
				}
			}
			if !yield(' ') {
				return
			}
		}
		for y := 1; y < yl; y++ {
			for x := xl - 1; x >= y; x-- {
				if !yield(runes[x][y+yl-1-x]) {
					return
				}
			}
			if !yield(' ') {
				return
			}
		}
	}
}

func diagonalCount(runes [][]rune) int {
	flattened := make([]rune, 0)
	for r := range diagonal(runes) {
		flattened = append(flattened, r)
	}
	return countXmas(string(flattened))
}

func countXmas(s string) int {
	return strings.Count(s, "XMAS") + strings.Count(s, "SAMX")
}

func part1(s string) int {
	lines := strings.Split(strings.ReplaceAll(s, "\r", ""), "\n")
	runes := runeSlices(lines)

	count := 0
	for _, line := range runes {
		count += countXmas(string(line))
	}
	count += diagonalCount(runes)

	t := rotate(runes)
	for _, line := range t {
		count += countXmas(string(line))
	}
	count += diagonalCount(t)

	return count
}

func part2(s string) int {
	lines := strings.Split(strings.ReplaceAll(s, "\r", ""), "\n")
	runes := runeSlices(lines)

	xl, yl := len(runes[0]), len(runes)
	kernel := make([][]rune, 3)
	for k := range kernel {
		kernel[k] = make([]rune, 3)
	}

	count := 0
	for y := range xl - 2 {
		for x := range yl - 2 {
			kernel[0] = runes[y][x : x+3]
			kernel[1] = runes[y+1][x : x+3]
			kernel[2] = runes[y+2][x : x+3]

			for range 4 {
				kernel = rotate(kernel)
				if kernel[0][0] == 'M' &&
					kernel[0][2] == 'M' &&
					kernel[1][1] == 'A' &&
					kernel[2][0] == 'S' &&
					kernel[2][2] == 'S' {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	path := os.Args[1]

	b, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("usage: day4 inputfile")
		os.Exit(1)
	}

	fmt.Println(part1(string(b)))
	fmt.Println(part2(string(b)))
}
