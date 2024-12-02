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

func totalDistance(a, b []int) int {
	sort.Ints(a)
	sort.Ints(b)
	abs := func(v lo.Tuple2[int, int]) int {
		if v.A > v.B {
			return v.A - v.B
		} else {
			return v.B - v.A
		}
	}
	return lo.SumBy(lo.Zip2(a, b), abs)
}

func similarityScore(a, b []int) int {
	occurances := func(v int, i int) lo.Tuple2[int, int] {
		return lo.T2(v, lo.Count(b, v))
	}
	multiply := func(t lo.Tuple2[int, int]) int {
		return t.A * t.B
	}
	return lo.SumBy(lo.Map(a, occurances), multiply)
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
