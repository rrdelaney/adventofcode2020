package main

import (
	"bufio"
	"fmt"
	"os"
)

type Slope struct {
	X, Y int
}

func numTreesHit(patterns [][]bool, slope Slope, c chan int) {
	treesHit := 0
	x := 0
	y := 0

	for y < len(patterns) {
		pattern := patterns[y]
		isHit := pattern[x%len(pattern)]
		if isHit {
			treesHit++
		}

		x += slope.X
		y += slope.Y
	}
	c <- treesHit
}

func main() {
	patterns := [][]bool{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		pattern := []bool{}
		for _, c := range scanner.Text() {
			pattern = append(pattern, string(c) == "#")
		}

		patterns = append(patterns, pattern)
	}

	slopes := []Slope{Slope{1, 1}, Slope{3, 1}, Slope{5, 1}, Slope{7, 1}, Slope{1, 2}}
	c := make(chan int)

	for _, s := range slopes {
		go numTreesHit(patterns, s, c)
	}

	numResults := 0
	totalTreeMult := 1
	for numResults < len(slopes) {
		select {
		case r := <-c:
			numResults++
			totalTreeMult *= r
		}
	}

	fmt.Println(totalTreeMult)
}
