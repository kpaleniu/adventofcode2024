package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"

	"github.com/samber/lo"
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
	return lo.Sum(distances(a, b))
}

func parse(r io.Reader) ([]int, []int, error) {
	a, b := make([]int, 0), make([]int, 0)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s := lo.Words(scanner.Text())
		if u, err := strconv.Atoi(s[0]); err == nil {
			a = append(a, u)
		}
		if v, err := strconv.Atoi(s[len(s)-1]); err == nil {
			b = append(b, v)
		}
	}
	if scanner.Err() != nil {
		return nil, nil, scanner.Err()
	}
	return a, b, nil
}

func countValuesOfIn(a, b []int) map[int]int {
	counter := make(map[int]int)
	for _, u := range a {
		for _, v := range b {
			if u == v {
				if i, ok := counter[u]; ok {
					counter[u] = i + 1
				} else {
					counter[u] = 1
				}
			}
		}
	}
	return counter
}

func similarityScore(a, b []int) int {
	return lo.Sum(lo.MapToSlice(countValuesOfIn(a, b), func(k int, v int) int { return k * v }))
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
	fmt.Println(similarityScore(a, b))
}
