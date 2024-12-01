package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func distances(a, b []int) []int {
	sort.Ints(a)
	sort.Ints(b)
	result := make([]int, min(len(a), len(b)))
	for i := range min(len(a), len(b)) {
		ma, mi := max(a[i], b[i]), min(a[i], b[i])
		result[i] = ma - mi
	}
	return result
}

func totalDistance(a, b []int) int {
	distance := 0
	for _, v := range distances(a, b) {
		distance += v
	}
	return distance
}

func parse(r io.Reader) ([]int, []int, error) {
	a, b := make([]int, 0), make([]int, 0)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		if u, err := strconv.Atoi(strings.TrimSpace(s[0])); err == nil {
			a = append(a, u)
		}
		if v, err := strconv.Atoi(strings.TrimSpace(s[len(s)-1])); err == nil {
			b = append(b, v)
		}
	}
	if scanner.Err() != nil {
		return nil, nil, scanner.Err()
	}
	return a, b, nil
}

func main() {
	path := os.Args[1]
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("usage: day1 inputfile")
		os.Exit(1)
	}
	defer f.Close()

	a, b, err := parse(f)
	if err != nil {
		fmt.Println("fail: could not parse input")
		os.Exit(1)
	}

	fmt.Println(totalDistance(a, b))
}
