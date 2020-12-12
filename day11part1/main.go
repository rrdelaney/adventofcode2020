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

	for y := posY - 1; y <= posY+1; y++ {
		if y < 0 || y >= len(seats) {
			continue
		}

		for x := posX - 1; x <= posX+1; x++ {
			if x < 0 || x >= len(seats[0]) {
				continue
			}
			if x == posX && y == posY {
				continue
			}

			seat := seats[y][x]
			switch seat {
			case 'L':
				empty++
			case '#':
				occupied++
			}
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
				if occupied >= 4 {
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
