package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Rule struct {
	Min, Max int
	Letter   string
}

type PasswordWithRule struct {
	Password string
	Rule     Rule
}

func CheckPassword(pw PasswordWithRule) int {
	seen := 0
	for _, c := range pw.Password {
		if pw.Rule.Letter == string(c) {
			seen++
		}
	}

	if pw.Rule.Min <= seen && seen <= pw.Rule.Max {
		return 1
	}

	return 0
}

func main() {
	pws := []PasswordWithRule{}

	scanner := bufio.NewScanner(os.Stdin)
	re := regexp.MustCompile(`^(\d+)\-(\d+) (\w): (\w+)$`)
	for scanner.Scan() {
		matches := re.FindStringSubmatch(scanner.Text())
		min, err := strconv.Atoi(matches[1])
		if err != nil {
			fmt.Printf("could not convert number: %v", err)
			os.Exit(1)
		}
		max, err := strconv.Atoi(matches[2])
		if err != nil {
			fmt.Printf("could not convert number: %v", err)
			os.Exit(1)
		}

		rule := Rule{min, max, matches[3]}
		pws = append(pws, PasswordWithRule{matches[4], rule})
	}

	valid := 0
	for _, pw := range pws {
		valid += CheckPassword(pw)
	}

	fmt.Println(valid)
}
