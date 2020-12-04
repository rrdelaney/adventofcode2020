package main

import (
	"bufio"
	"fmt"
	"os"
)

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

	treesHit := 0
	x := 0
	for _, pattern := range patterns {
		hit := pattern[x%len(pattern)]
		if hit {
			treesHit++
		}

		x += 3
	}

	fmt.Println(treesHit)
}
