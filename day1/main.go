// Solution for Advent of Code 2020 -- Day 1
// https://adventofcode.com/2020/day/1

package main

import (
	"fmt"

	"github.com/huderlem/adventofcode2020/util"
)

func part1() int {
	expenses := util.ReadFileInts("input.txt")
	for i := 0; i < len(expenses)-2; i++ {
		for j := i + 1; j < len(expenses)-1; j++ {
			sum := expenses[i] + expenses[j]
			if sum == 2020 {
				return expenses[i] * expenses[j]
			}
		}
	}
	// Valid expense combination not found.
	return -1
}

func part2() int {
	expenses := util.ReadFileInts("input.txt")
	for i := 0; i < len(expenses)-2; i++ {
		for j := i + 1; j < len(expenses)-1; j++ {
			for k := j + 1; k < len(expenses); k++ {
				sum := expenses[i] + expenses[j] + expenses[k]
				if sum == 2020 {
					return expenses[i] * expenses[j] * expenses[k]
				}
			}
		}
	}
	// Valid expense combination not found.
	return -1
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
