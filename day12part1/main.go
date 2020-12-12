package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	dir := 90
	x := 0
	y := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		action := line[:1]
		arg, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		switch action {
		case "N":
			y += arg
		case "S":
			y -= arg
		case "E":
			x += arg
		case "W":
			x -= arg
		case "R":
			dir = (dir + 360 + arg) % 360
		case "L":
			dir = (dir + 360 - arg) % 360
		case "F":
			switch dir {
			case 0:
				y += arg
			case 90:
				x += arg
			case 180:
				y -= arg
			case 270:
				x -= arg
			}
		default:
		}
	}

	fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)))
}
