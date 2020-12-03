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
	pMin := string(pw.Password[pw.Rule.Min-1]) == pw.Rule.Letter
	pMax := string(pw.Password[pw.Rule.Max-1]) == pw.Rule.Letter

	if pMin != pMax {
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
