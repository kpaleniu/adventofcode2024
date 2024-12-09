package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var pattern = regexp.MustCompile(`(do\(\)|don't\(\)|mul\((\d+),(\d+)\))`)

func part1(s string) int {
	sum := 0
	for _, m := range pattern.FindAllStringSubmatch(s, -1) {
		u, _ := strconv.Atoi(m[len(m)-2])
		v, _ := strconv.Atoi(m[len(m)-1])
		sum += u * v
	}
	return sum
}

func part2(s string) int {
	sum := 0
	ignore := false
	for _, m := range pattern.FindAllStringSubmatch(s, -1) {
		if m[0] == "don't()" {
			ignore = true
		}
		if m[0] == "do()" {
			ignore = false
		}
		if ignore {
			continue
		}
		u, _ := strconv.Atoi(m[len(m)-2])
		v, _ := strconv.Atoi(m[len(m)-1])
		sum += u * v
	}
	return sum
}

func main() {
	path := os.Args[1]

	b, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("usage: day3 inputfile")
		os.Exit(1)
	}

	fmt.Println(part1(string(b)))
	fmt.Println(part2(string(b)))
}
