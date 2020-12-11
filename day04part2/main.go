package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	Byr, Iyr, Eyr, Hgt, Hcl, Ecl, Pid, Cid string
}

func (p Passport) IsValid() bool {
	byrNum, err := strconv.Atoi(p.Byr)
	if err != nil {
		return false
	}
	if !(1920 <= byrNum && byrNum <= 2002) {
		return false
	}

	iyrNum, err := strconv.Atoi(p.Iyr)
	if err != nil {
		return false
	}
	if !(2010 <= iyrNum && iyrNum <= 2020) {
		return false
	}

	eyrNum, err := strconv.Atoi(p.Eyr)
	if err != nil {
		return false
	}
	if !(2020 <= eyrNum && eyrNum <= 2030) {
		return false
	}

	if strings.HasSuffix(p.Hgt, "cm") {
		height, err := strconv.Atoi(strings.TrimRight(p.Hgt, "cm"))
		if err != nil {
			return false
		}
		if !(150 <= height && height <= 193) {
			return false
		}
	} else if strings.HasSuffix(p.Hgt, "in") {
		height, err := strconv.Atoi(strings.TrimRight(p.Hgt, "in"))
		if err != nil {
			return false
		}
		if !(59 <= height && height <= 76) {
			return false
		}
	} else {
		return false
	}

	hclRe := regexp.MustCompile(`^#[a-f0-9]{6}$`)
	if !hclRe.MatchString(p.Hcl) {
		return false
	}

	eclRe := regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	if !eclRe.MatchString(p.Ecl) {
		return false
	}

	pidRe := regexp.MustCompile(`^\d{9}$`)
	if !pidRe.MatchString(p.Pid) {
		return false
	}

	return true
}

func PassportFromString(s string) Passport {
	pp := Passport{}

	for _, match := range strings.Split(s, " ") {
		keyVal := strings.Split(match, ":")
		switch keyVal[0] {
		case "byr":
			pp.Byr = keyVal[1]
		case "iyr":
			pp.Iyr = keyVal[1]
		case "eyr":
			pp.Eyr = keyVal[1]
		case "hgt":
			pp.Hgt = keyVal[1]
		case "hcl":
			pp.Hcl = keyVal[1]
		case "ecl":
			pp.Ecl = keyVal[1]
		case "pid":
			pp.Pid = keyVal[1]
		case "cid":
			pp.Cid = keyVal[1]
		}
	}

	return pp
}

func main() {
	pps := []Passport{}

	scanner := bufio.NewScanner(os.Stdin)

	buf := ""
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			buf += line + " "
			continue
		}

		pps = append(pps, PassportFromString(buf))
		buf = ""
	}
	pps = append(pps, PassportFromString(buf))

	numValid := 0
	for _, pp := range pps {
		if pp.IsValid() {
			numValid++
		}
	}

	fmt.Println(numValid)
}
