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

func AllYesAnswersFromStrings(ss []string) map[rune]bool {
	allYesSoFar := AnswersFromString(ss[0])
	for _, s := range ss[1:] {
		currAns := AnswersFromString(s)
		for k := range allYesSoFar {
			if currAns[k] != true {
				delete(allYesSoFar, k)
			}
		}
	}

	return allYesSoFar
}

func main() {
	answers := []map[rune]bool{}

	scanner := bufio.NewScanner(os.Stdin)
	buf := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			buf = append(buf, line)
			continue
		}

		answers = append(answers, AllYesAnswersFromStrings(buf))
		buf = nil
	}
	answers = append(answers, AllYesAnswersFromStrings(buf))

	numAnswers := 0
	for _, ans := range answers {
		numAnswers += len(ans)
	}

	fmt.Println(numAnswers)
}
