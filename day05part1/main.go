package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func seatIDFromInstructions(ins string) float64 {
	front := 0.0
	back := 127.0
	left := 0.0
	right := 7.0

	for _, in := range ins {
		switch string(in) {
		case "B":
			front = math.Ceil((front + back) / 2)
		case "F":
			back = math.Floor((front + back) / 2)
		case "L":
			right = math.Floor((left + right) / 2)
		case "R":
			left = math.Ceil((left + right) / 2)
		}
	}

	return (front * 8) + left
}

func main() {
	ins := []string{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ins = append(ins, scanner.Text())
	}

	maxSeatID := 0.0
	for _, in := range ins {
		maxSeatID = math.Max(maxSeatID, seatIDFromInstructions(in))
	}

	fmt.Println(maxSeatID)
}
