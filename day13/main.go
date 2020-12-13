// Solution for Advent of Code 2020 -- Day 13
// https://adventofcode.com/2020/day/13

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/huderlem/adventofcode2020/util"
)

func parseInput() (int, []int) {
	lines := util.ReadFileLines("input.txt")
	startTime, _ := strconv.Atoi(lines[0])
	busIDs := []int{}
	for _, rawID := range strings.Split(lines[1], ",") {
		if rawID != "x" {
			id, _ := strconv.Atoi(rawID)
			busIDs = append(busIDs, id)
		} else {
			busIDs = append(busIDs, -1)
		}
	}
	return startTime, busIDs
}

func part1() int {
	startTime, busIDs := parseInput()
	minTime := math.MaxInt32
	minID := busIDs[0]
	for _, id := range busIDs {
		if id == -1 {
			continue
		}
		mod := (startTime % id)
		if mod == 0 {
			mod = id
		}
		time := startTime + (id - mod)
		if time < minTime {
			minTime = time
			minID = id
		}
	}
	return (minTime - startTime) * minID
}

func part2() int {
	_, busIDs := parseInput()
	result := 0
	step := 1
	// Chinese remainder theorem... didn't know this existed.
	// I attempted a solution using extended gcd, which passed the test cases
	// but failed the real input and gave up.
	for i, id := range busIDs {
		if id < 1 {
			continue
		}
		for (result+i)%id != 0 {
			result += step
		}
		step *= id
	}
	return result
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
