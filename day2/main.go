// Solution for Advent of Code 2020 -- Day 2
// https://adventofcode.com/2020/day/2

package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/huderlem/adventofcode2020/util"
)

func part1() int {
	items := util.ReadFileLines("input.txt")
	numValid := 0
	re, _ := regexp.Compile(`(?P<min>\d+)-(?P<max>\d+)\s+(?P<letter>\w):\s+(?P<password>\w+)`)
	for _, item := range items {
		match := re.FindStringSubmatch(item)
		result := make(map[string]string)
		for i, name := range re.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = match[i]
			}
		}
		min, _ := strconv.Atoi(result["min"])
		max, _ := strconv.Atoi(result["max"])
		letter := result["letter"]
		password := result["password"]
		numLetters := strings.Count(password, letter)
		if numLetters >= min && numLetters <= max {
			numValid++
		}
	}
	return numValid
}

func part2() int {
	items := util.ReadFileLines("input.txt")
	numValid := 0
	re, _ := regexp.Compile(`(?P<pos1>\d+)-(?P<pos2>\d+)\s+(?P<letter>\w):\s+(?P<password>\w+)`)
	for _, item := range items {
		match := re.FindStringSubmatch(item)
		result := make(map[string]string)
		for i, name := range re.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = match[i]
			}
		}
		pos1, _ := strconv.Atoi(result["pos1"])
		pos2, _ := strconv.Atoi(result["pos2"])
		pos1--
		pos2--
		letter := result["letter"]
		password := result["password"]
		if pos1 > len(password) || pos2 > len(password) {
			continue
		}
		validPositions := 0
		if string(password[pos1]) == letter {
			validPositions++
		}
		if string(password[pos2]) == letter {
			validPositions++
		}
		if validPositions == 1 {
			numValid++
		}
	}
	return numValid
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
