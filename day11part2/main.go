package main

import (
	"bufio"
	"fmt"
	"os"
)

func TotalOccupied(seats [][]rune) int {
	occupied := 0

	for _, row := range seats {
		for _, seat := range row {
			if seat == '#' {
				occupied += 1
			}
		}
	}

	return occupied
}

func NumAdjacentSeats(posX, posY int, seats [][]rune) (int, int) {
	empty := 0
	occupied := 0

	directions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 0}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, direction := range directions {
		if direction[0] == 0 && direction[1] == 0 {
			continue
		}

		x := posX + direction[0]
		y := posY + direction[1]
		for 0 <= y && y < len(seats) && 0 <= x && x < len(seats[0]) {
			switch seats[y][x] {
			case '.':
				x += direction[0]
				y += direction[1]
				continue
			case 'L':
				empty++
			case '#':
				occupied++
			}
			break
		}
	}

	return empty, occupied
}

func Step(seats [][]rune) [][]rune {
	nextSeats := [][]rune{}

	for posY, row := range seats {
		nextRow := []rune{}

		for posX, seat := range row {
			nextSeat := seat
			_, occupied := NumAdjacentSeats(posX, posY, seats)

			switch seat {
			case 'L':
				if occupied == 0 {
					nextSeat = '#'
				}
			case '#':
				if occupied >= 5 {
					nextSeat = 'L'
				}
			}

			nextRow = append(nextRow, nextSeat)
		}

		nextSeats = append(nextSeats, nextRow)
	}

	return nextSeats
}

func Serialize(seats [][]rune) string {
	s := ""

	for _, row := range seats {
		for _, seat := range row {
			s += string(seat)
		}
		s += "\n"
	}

	return s
}

func main() {
	seats := [][]rune{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		seats = append(seats, []rune(scanner.Text()))
	}

	for {
		nextSeats := Step(seats)
		if Serialize(nextSeats) == Serialize(seats) {
			break
		}

		seats = nextSeats
	}

	fmt.Println(TotalOccupied(seats))
}
