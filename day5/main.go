// Solution for Advent of Code 2020 -- Day 5
// https://adventofcode.com/2020/day/5

package main

import (
	"fmt"

	"github.com/huderlem/adventofcode2020/util"
)

// The binary partitioning is binary representation of the number. So, we
// can use bitwise operations to construct the value, setting each bit
// individually.
func resolvePartitioning(row string, bitIndicator rune) int {
	result := 0
	for i, v := range row {
		if v == bitIndicator {
			bitIndex := len(row) - i - 1
			result |= 1 << bitIndex
		}
	}
	return result
}

func part1() int {
	boardingPasses := util.ReadFileLines("input.txt")
	maxSeatID := -1
	for _, boardingPass := range boardingPasses {
		row := resolvePartitioning(boardingPass[:7], 'B')
		column := resolvePartitioning(boardingPass[7:], 'R')
		seatID := row*8 + column
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}
	return maxSeatID
}

func part2() int {
	boardingPasses := util.ReadFileLines("input.txt")
	filledSeats := make([]bool, 128*8)
	for _, boardingPass := range boardingPasses {
		row := resolvePartitioning(boardingPass[:7], 'B')
		column := resolvePartitioning(boardingPass[7:], 'R')
		seatID := row*8 + column
		filledSeats[seatID] = true
	}
	// Skip over the unfilled seats at the front of the plane.
	seatID := 0
	for !filledSeats[seatID] {
		seatID++
	}
	// Then, find the first unfilled seat.
	for filledSeats[seatID] {
		seatID++
	}
	return seatID
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
