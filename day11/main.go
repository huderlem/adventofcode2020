// Solution for Advent of Code 2020 -- Day 11
// https://adventofcode.com/2020/day/11

package main

import (
	"fmt"

	"github.com/huderlem/adventofcode2020/util"
)

type seatStatus int
type seatsLayout [][]seatStatus

const (
	missing seatStatus = iota
	empty
	occupied
)

type occupiedCounter func(x, y int, seats seatsLayout) int

func (seats seatsLayout) getDimensions() (int, int) {
	return len(seats), len(seats[0])
}

func parseSeatGrid() seatsLayout {
	lines := util.ReadFileLines("input.txt")
	width := len(lines[0])
	grid := make(seatsLayout, width)
	for _, line := range lines {
		for x, cell := range line {
			seat := missing
			if byte(cell) == 'L' {
				seat = empty
			} else if byte(cell) == '#' {
				seat = occupied
			}
			grid[x] = append(grid[x], seat)
		}
	}
	return grid
}

func getOccupiedNeighborCount(x, y int, seats seatsLayout) int {
	width, height := seats.getDimensions()
	count := 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i < 0 || i >= width || j < 0 || j >= height || (i == x && j == y) {
				continue
			}
			if seats[i][j] == occupied {
				count++
			}
		}
	}
	return count
}

func findSeatInDirection(x, y, dx, dy int, seats seatsLayout) seatStatus {
	width, height := seats.getDimensions()
	i := x + dx
	j := y + dy
	for i >= 0 && i <= width-1 && j >= 0 && j <= height-1 {
		seat := seats[i][j]
		if seat != missing {
			return seat
		}
		i += dx
		j += dy
	}
	return missing
}

func getVisibleNeighborCount(x, y int, seats seatsLayout) int {
	count := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			if findSeatInDirection(x, y, dx, dy, seats) == occupied {
				count++
			}
		}
	}
	return count
}

func seatsAreEqual(a, b seatsLayout) bool {
	width, height := a.getDimensions()
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if a[x][y] != b[x][y] {
				return false
			}
		}
	}
	return true
}

func applySeatingRules(seats seatsLayout, counterFunc occupiedCounter, occupiedLimit int) (seatsLayout, bool) {
	changed := false
	width, height := seats.getDimensions()
	newSeats := make(seatsLayout, width)
	for i := 0; i < width; i++ {
		newSeats[i] = make([]seatStatus, height)
		copy(newSeats[i], seats[i])
	}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			seat := seats[x][y]
			if seat == missing {
				continue
			}
			occupiedCount := counterFunc(x, y, seats)
			if seat == empty && occupiedCount == 0 {
				if newSeats[x][y] != occupied {
					changed = true
				}
				newSeats[x][y] = occupied
			} else if seat == occupied && occupiedCount >= occupiedLimit {
				if newSeats[x][y] != empty {
					changed = true
				}
				newSeats[x][y] = empty
			}
		}
	}
	return newSeats, changed
}

func countOccupiedSeats(seats seatsLayout) int {
	width, height := seats.getDimensions()
	count := 0
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if seats[x][y] == occupied {
				count++
			}
		}
	}
	return count
}

func part1() int {
	seats := parseSeatGrid()
	for {
		newSeats, changed := applySeatingRules(seats, getOccupiedNeighborCount, 4)
		if !changed {
			return countOccupiedSeats(newSeats)
		}
		seats = newSeats
	}
}

func part2() int {
	seats := parseSeatGrid()
	for {
		newSeats, changed := applySeatingRules(seats, getVisibleNeighborCount, 5)
		if !changed {
			return countOccupiedSeats(newSeats)
		}
		seats = newSeats
	}
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
