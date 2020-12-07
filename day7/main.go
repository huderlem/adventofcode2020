// Solution for Advent of Code 2020 -- Day 7
// https://adventofcode.com/2020/day/7

package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/huderlem/adventofcode2020/util"
)

type bagRules map[string]map[string]int

func parseBagRules() bagRules {
	rules := map[string]map[string]int{}
	lines := util.ReadFileLines("input.txt")
	mainBagRe, _ := regexp.Compile(`^(?P<container>\w+ \w+) bags contain (?P<rest>.+)`)
	rulesRe, _ := regexp.Compile(`(?P<num>\d+) (?P<color>\w+ \w+) bags?`)
	for _, line := range lines {
		match := mainBagRe.FindStringSubmatch(line)
		mainMatches := make(map[string]string)
		for i, name := range mainBagRe.SubexpNames() {
			if i != 0 && name != "" {
				mainMatches[name] = match[i]
			}
		}
		containerBag := mainMatches["container"]
		if _, ok := rules[containerBag]; ok {
			panic(fmt.Sprintf("Multiple rules defined for %s", containerBag))
		}
		rules[containerBag] = map[string]int{}
		parts := strings.Split(mainMatches["rest"], ",")
		for _, part := range parts {
			matches := rulesRe.FindStringSubmatch(part)
			if len(matches) == 0 {
				continue
			}
			partMatches := make(map[string]string)
			for i, name := range rulesRe.SubexpNames() {
				if i != 0 && name != "" {
					partMatches[name] = matches[i]
				}
			}
			numBags, _ := strconv.Atoi(partMatches["num"])
			rules[containerBag][partMatches["color"]] = numBags
		}
	}
	return rules
}

func bagHoldsColor(color string, holds map[string]int, rules bagRules, visited map[string]bool) (bool, map[string]bool) {
	for holdColor, _ := range holds {
		if holdColor == color {
			return true, visited
		}
	}
	for holdColor, _ := range holds {
		if _, ok := visited[holdColor]; ok {
			continue
		}
		var holdsColor bool
		holdsColor, visited = bagHoldsColor(color, rules[holdColor], rules, visited)
		if holdsColor {
			return true, visited
		}
		visited[holdColor] = true
	}
	return false, visited
}

func countHeldBags(color string, rules bagRules) int {
	if len(rules[color]) == 0 {
		return 0
	}
	total := 0
	for holdColor, count := range rules[color] {
		total += count * (1 + countHeldBags(holdColor, rules))
	}
	return total
}

func part1() int {
	rules := parseBagRules()
	num := 0
	for _, holds := range rules {
		visited := map[string]bool{}
		holdsColor, _ := bagHoldsColor("shiny gold", holds, rules, visited)
		if holdsColor {
			num++
		}
	}
	return num
}

func part2() int {
	rules := parseBagRules()
	return countHeldBags("shiny gold", rules)
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
