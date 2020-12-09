package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	Operation string
	Argument  int
}

func InstructionFromString(s string) (Instruction, error) {
	parts := strings.Split(s, " ")
	ins := Instruction{Operation: parts[0]}
	arg, err := strconv.Atoi(parts[1])
	if err != nil {
		return ins, err
	}

	ins.Argument = arg
	return ins, nil
}

type Console struct {
	Program     []Instruction
	Position    int
	Accumulator int
}

func (c *Console) Step() {
	ins := c.Program[c.Position]
	switch ins.Operation {
	case "nop":
		c.Position++
	case "acc":
		c.Accumulator += ins.Argument
		c.Position++
	case "jmp":
		c.Position += ins.Argument
	}
}

func main() {
	c := &Console{[]Instruction{}, 0, 0}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ins, err := InstructionFromString(scanner.Text())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		c.Program = append(c.Program, ins)
	}

	seenPositions := map[int]bool{}
	for {
		if seenPositions[c.Position] {
			break
		}

		seenPositions[c.Position] = true
		c.Step()
	}

	fmt.Println(c.Accumulator)
}
