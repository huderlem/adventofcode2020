// Solution for Advent of Code 2020 -- Day 15
// https://adventofcode.com/2020/day/15

package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/huderlem/adventofcode2020/util"
)

func parseInput() []int {
	input := util.ReadFileString("input.txt")
	nums := []int{}
	for _, n := range strings.Split(input, ",") {
		value, _ := strconv.Atoi(n)
		nums = append(nums, value)
	}
	return nums
}

func playMemoryGame(nums []int, n int) int {
	prevIndexes := map[int]int{}
	for i, num := range nums {
		prevIndexes[num] = i
	}
	for i := len(nums); i < n; i++ {
		lastNum := nums[len(nums)-1]
		if index, ok := prevIndexes[lastNum]; ok && index != i-1 {
			nums = append(nums, i-prevIndexes[lastNum]-1)
		} else {
			nums = append(nums, 0)
		}
		prevIndexes[lastNum] = i - 1
	}

	return nums[len(nums)-1]
}

func part1() int {
	nums := parseInput()
	return playMemoryGame(nums, 2020)
}

func part2() int {
	nums := parseInput()
	return playMemoryGame(nums, 30000000)
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
