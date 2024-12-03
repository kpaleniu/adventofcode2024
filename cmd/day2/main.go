package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func isSafe(levels []int) bool {
	if len(levels) == 0 {
		return true
	}
	inc := levels[len(levels)-1]-levels[0] > 0
	for i, v := range levels {
		p := levels[max(i-1, 0)]
		// increasing and decreasing
		if inc && v < p {
			return false
		}
		if !inc && v > p {
			return false
		}
		// not increasing nor decreasing
		if i > 0 && v == p {
			return false
		}
		// too large increase or decrease
		if v-p > 3 || v-p < -3 {
			return false
		}
	}
	return true
}

func levelIterator(r io.Reader) func(func([]int) bool) {
	return func(yield func(levels []int) bool) {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			words := lo.Words(scanner.Text())
			levels := lo.FilterMap(words, func(s string, i int) (int, bool) {
				v, err := strconv.Atoi(s)
				return v, err == nil
			})
			if !yield(levels) {
				return
			}
		}
	}
}

func part1(r io.Reader) int {
	safe := 0
	for levels := range levelIterator(r) {
		if isSafe(levels) {
			safe += 1
		}
	}
	return safe
}

func part2(r io.Reader) int {
	safe := 0
	for levels := range levelIterator(r) {
		s := isSafe(levels)
		if !s {
			// brute-force attempt to remove one element if it makes the level safe
			for i := range len(levels) {
				u := slices.Concat(levels[:i], levels[i+1:])
				if isSafe(u) {
					s = true
					break
				}
			}
		}
		if s {
			safe += 1
		}
	}
	return safe
}

func main() {
	path := os.Args[1]

	b, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("usage: day2 inputfile")
		os.Exit(1)
	}

	fmt.Println(part1(strings.NewReader(string(b))))
	fmt.Println(part2(strings.NewReader(string(b))))
}
