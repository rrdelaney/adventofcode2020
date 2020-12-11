package main

import (
	"bufio"
	"fmt"
	"os"
)

func AnswersFromString(s string) map[rune]bool {
	m := map[rune]bool{}
	for _, r := range s {
		m[r] = true
	}

	return m
}

func main() {
	answers := []map[rune]bool{}

	scanner := bufio.NewScanner(os.Stdin)
	buf := ""
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			buf += line
			continue
		}

		answers = append(answers, AnswersFromString(buf))
		buf = ""
	}
	answers = append(answers, AnswersFromString(buf))

	numAnswers := 0
	for _, ans := range answers {
		numAnswers += len(ans)
	}

	fmt.Println(numAnswers)
}
