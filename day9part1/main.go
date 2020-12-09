package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const preambleSize = 25

func Pop(q []int, size int) []int {
	diff := len(q) - size
	if diff <= 0 {
		return q
	}

	return q[diff:]
}

func Valid(q []int, next int, size int) bool {
	if len(q) < size {
		return true
	}

	for index, i := range q {
		for _, j := range q[index:] {
			if i+j == next {
				return true
			}
		}
	}
	return false
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

		if !Valid(q, next, preambleSize) {
			fmt.Println(next)
			os.Exit(0)
		}

		q = append(q, next)
		q = Pop(q, preambleSize)
	}
}
