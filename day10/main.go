// Solution for Advent of Code 2020 -- Day 10
// https://adventofcode.com/2020/day/10

package main

import (
	"fmt"
	"sort"

	"github.com/huderlem/adventofcode2020/util"
)

func getSortedJoltages() []int {
	joltages := util.ReadFileInts("input.txt")
	joltages = append(joltages, 0)
	sort.Ints(joltages)
	deviceRating := joltages[len(joltages)-1] + 3
	joltages = append(joltages, deviceRating)
	return joltages
}

func getJoltageOptions(joltages []int) []int {
	options := make([]int, len(joltages))
	for i := 0; i < len(joltages); i++ {
		for j := 1; j < 4; j++ {
			if i+j < len(joltages) && joltages[i+j]-joltages[i] <= 3 {
				options[i]++
			}
		}
	}
	return options
}

func part1() int {
	joltages := getSortedJoltages()
	diffs := make([]int, 4)
	for i := 1; i < len(joltages); i++ {
		diff := joltages[i] - joltages[i-1]
		diffs[diff]++
	}
	return diffs[1] * diffs[3]
}

func part2() int {
	joltages := getSortedJoltages()
	options := getJoltageOptions(joltages)
	// Recursion seems like a natural solution at first, but it would probably
	// collapse due to the large number of branches and stack limit. Instead,
	// the result can be calculated using a dynamic programming approach,
	// summing the cummulative number of combinations starting at the back of
	// the list.
	counts := make([]int, len(options))
	counts[len(options)-1] = 1
	for i := len(options) - 2; i >= 0; i-- {
		for j := 1; j <= options[i]; j++ {
			counts[i] += counts[i+j]
		}
	}
	return counts[0]
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
