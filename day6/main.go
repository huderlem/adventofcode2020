// Solution for Advent of Code 2020 -- Day 6
// https://adventofcode.com/2020/day/6

package main

import (
	"fmt"

	"github.com/huderlem/adventofcode2020/util"
)

func parseDeclarations() [][]string {
	chunks := util.ReadFileChunks("input.txt")
	declarations := [][]string{}
	for _, chunk := range chunks {
		curDeclaration := []string{}
		for _, line := range chunk {
			curDeclaration = append(curDeclaration, line)
		}
		declarations = append(declarations, curDeclaration)
	}

	return declarations
}

func part1() int {
	declarations := parseDeclarations()
	sum := 0
	for _, group := range declarations {
		uniqueDeclarations := map[rune]struct{}{}
		for _, answers := range group {
			for _, letter := range answers {
				uniqueDeclarations[letter] = struct{}{}
			}
		}
		sum += len(uniqueDeclarations)
	}

	return sum
}

func part2() int {
	declarations := parseDeclarations()
	sum := 0
	for _, group := range declarations {
		declarationCounts := map[rune]int{}
		for _, answers := range group {
			for _, letter := range answers {
				if _, ok := declarationCounts[letter]; !ok {
					declarationCounts[letter] = 0
				}
				declarationCounts[letter]++
			}
		}
		numAllYes := 0
		for _, num := range declarationCounts {
			if num == len(group) {
				numAllYes++
			}
		}
		sum += numAllYes
	}

	return sum
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
