package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const preambleSize = 25

func Sum(q []int) int {
	t := 0
	for _, n := range q {
		t += n
	}
	return t
}

func Min(q []int) int {
	m := math.MaxInt64
	for _, n := range q {
		if n < m {
			m = n
		}
	}

	return m
}

func Max(q []int) int {
	m := math.MinInt64
	for _, n := range q {
		if n > m {
			m = n
		}
	}

	return m
}

func Valid(q []int, pos int) bool {
	if pos <= preambleSize {
		return true
	}

	next := q[pos]
	scope := q[pos-preambleSize : pos]
	for index, i := range scope {
		for _, j := range scope[index:] {
			if i+j == next {
				return true
			}
		}
	}

	return false
}

func FindInvalid(q []int) (int, error) {
	for i, val := range q {
		if !Valid(q, i) {
			return val, nil
		}
	}

	return 0, fmt.Errorf("no invalid value")
}

func FindRangeSummingTo(q []int, val int) []int {
	for low := range q {
		for high := range q[low:] {
			r := q[low:(low + high)]
			if Sum(r) == val {
				return r
			}
		}
	}

	return nil
}

func main() {
	q := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		next, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		q = append(q, next)
	}

	invalid, err := FindInvalid(q)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	invalidRange := FindRangeSummingTo(q, invalid)
	fmt.Println(Min(invalidRange) + Max(invalidRange))
}
