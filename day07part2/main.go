package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Rule struct {
	Color  string
	Number int
}

func ColorAndRulesFromString(s string) (string, []Rule, error) {
	colorRe := regexp.MustCompile(`(\w[\w ]*) bags? contain`)
	maybeColor := colorRe.FindStringSubmatch(s)
	if len(maybeColor) != 2 {
		return "", nil, fmt.Errorf("Could not find color in: %v", s)
	}
	color := maybeColor[1]

	containsRe := regexp.MustCompile(`(\d+) (\w[\w ]*) bags?`)
	maybeRules := containsRe.FindAllStringSubmatch(s, -1)
	rules := []Rule{}
	for _, match := range maybeRules {
		if len(match) != 3 {
			return "", nil, fmt.Errorf("Could not find rule in: %v", match)
		}
		matchColor := match[2]
		matchNum, err := strconv.Atoi(match[1])
		if err != nil {
			return "", nil, err
		}

		rules = append(rules, Rule{matchColor, matchNum})
	}

	return color, rules, nil
}

func BagsWithin(ruleMap map[string][]Rule, color string) int {
	bagsWithin := 0
	for _, rule := range ruleMap[color] {
		bagsWithin += rule.Number + (rule.Number * BagsWithin(ruleMap, rule.Color))
	}

	return bagsWithin
}

func main() {
	rules := map[string][]Rule{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		color, rule, err := ColorAndRulesFromString(scanner.Text())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		rules[color] = rule
	}

	fmt.Println(BagsWithin(rules, "shiny gold"))
}
