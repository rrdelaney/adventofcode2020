package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func Rotate(x, y, deg int) (int, int) {
	s, c := math.Sincos(math.Pi * float64(deg) / 180.0)
	si, ci := int(s), int(c)
	return x*ci - y*si, x*si + y*ci
}

func main() {
	wx := 10
	wy := 1
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
			wy += arg
		case "S":
			wy -= arg
		case "E":
			wx += arg
		case "W":
			wx -= arg
		case "R":
			wx, wy = Rotate(wx, wy, -arg)
		case "L":
			wx, wy = Rotate(wx, wy, arg)
		case "F":
			x += wx * arg
			y += wy * arg
		}
	}

	fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)))
}
