package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Passport struct {
	Byr, Iyr, Eyr, Hgt, Hcl, Ecl, Pid, Cid string
}

func (p Passport) IsValid() bool {
	return p.Byr != "" && p.Iyr != "" && p.Eyr != "" && p.Hgt != "" && p.Hcl != "" && p.Ecl != "" && p.Pid != ""
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
