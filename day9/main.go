// Solution for Advent of Code 2020 -- Day 9
// https://adventofcode.com/2020/day/9

package main

import (
	"fmt"

	"github.com/huderlem/adventofcode2020/util"
)

func hasSumPair(value int, values []int) bool {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] != values[j] && values[i]+values[j] == value {
				return true
			}
		}
	}
	return false
}

func findSumSet(value int, values []int) (bool, []int) {
	for i := 0; i < len(values)-1; i++ {
		sum := 0
		j := i
		for ; j < len(values) && sum < value; j++ {
			sum += values[j]
		}
		if sum == value {
			return true, values[i:j]
		}
	}
	return false, []int{}
}

func getMinMax(values []int) (int, int) {
	if len(values) == 0 {
		return -1, -1
	}
	min := values[0]
	max := values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

func part1() int {
	nums := util.ReadFileInts("input.txt")
	for i := 25; i < len(nums); i++ {
		hasPair := hasSumPair(nums[i], nums[i-25:i])
		if !hasPair {
			return nums[i]
		}
	}
	return -1
}

func part2() int {
	nums := util.ReadFileInts("input.txt")
	invalidNumber := part1()
	found, sumSet := findSumSet(invalidNumber, nums)
	if !found {
		return -1
	}
	min, max := getMinMax(sumSet)
	return min + max
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
