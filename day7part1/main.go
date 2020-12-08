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

func CanContainShinyGold(ruleMap map[string][]Rule, color string, seen map[string]bool) bool {
	if seen[color] == true {
		return false
	}
	seen[color] = true

	rules := ruleMap[color]
	for _, rule := range rules {
		if rule.Color == "shiny gold" {
			return true
		}
		if CanContainShinyGold(ruleMap, rule.Color, seen) {
			return true
		}
	}

	return false
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

	numCanContainShinyGold := 0
	for color := range rules {
		if CanContainShinyGold(rules, color, make(map[string]bool)) {
			numCanContainShinyGold++
		}
	}

	fmt.Println(numCanContainShinyGold)
}
