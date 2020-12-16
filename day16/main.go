// Solution for Advent of Code 2020 -- Day 15
// https://adventofcode.com/2020/day/15

package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/huderlem/adventofcode2020/util"
)

type ticketRule struct {
	min, max int
}

func (t ticketRule) accepts(num int) bool {
	return num >= t.min && num <= t.max
}

func parseRules(lines []string) map[string][]ticketRule {
	rules := map[string][]ticketRule{}
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		rawRanges := strings.Split(parts[1], " or ")
		ranges := make([]ticketRule, len(rawRanges))
		for i, r := range rawRanges {
			vals := strings.Split(r, "-")
			minimum, _ := strconv.Atoi(vals[0])
			maximum, _ := strconv.Atoi(vals[1])
			ranges[i] = ticketRule{min: minimum, max: maximum}
		}
		rules[parts[0]] = ranges
	}
	return rules
}

func parseTickets(lines []string) [][]int {
	tickets := [][]int{}
	for _, line := range lines[1:] {
		ticket := []int{}
		nums := strings.Split(line, ",")
		for _, num := range nums {
			v, _ := strconv.Atoi(num)
			ticket = append(ticket, v)
		}
		tickets = append(tickets, ticket)
	}
	return tickets
}

func parseInput() (map[string][]ticketRule, []int, [][]int) {
	chunks := util.ReadFileChunks("input.txt")
	rules := parseRules(chunks[0])
	myTicket := parseTickets(chunks[1])[0]
	tickets := parseTickets(chunks[2])
	return rules, myTicket, tickets
}

func part1() int {
	rules, _, tickets := parseInput()
	invalidFields := []int{}
	for _, ticket := range tickets {
		for _, v := range ticket {
			valid := false
			for _, ranges := range rules {
				for _, r := range ranges {
					if r.accepts(v) {
						valid = true
					}
				}
			}
			if !valid {
				invalidFields = append(invalidFields, v)
			}
		}
	}

	sum := 0
	for _, field := range invalidFields {
		sum += field
	}
	return sum
}

func getValidTickets(tickets [][]int, rules map[string][]ticketRule) [][]int {
	validTickets := [][]int{}
	for _, ticket := range tickets {
		allValid := true
		for _, v := range ticket {
			valid := false
			for _, ranges := range rules {
				for _, r := range ranges {
					if r.accepts(v) {
						valid = true
					}
				}
			}
			if !valid {
				allValid = false
			}
		}
		if allValid {
			validTickets = append(validTickets, ticket)
		}
	}
	return validTickets
}

func initPossibilities(tickets [][]int, rules map[string][]ticketRule) []map[string]struct{} {
	possibilities := []map[string]struct{}{}
	for i := 0; i < len(tickets[0]); i++ {
		p := map[string]struct{}{}
		for ruleName, _ := range rules {
			p[ruleName] = struct{}{}
		}
		possibilities = append(possibilities, p)
	}
	return possibilities
}

func narrowPossibilities(tickets [][]int, possibilities []map[string]struct{}, rules map[string][]ticketRule) []map[string]struct{} {
	for _, ticket := range tickets {
		for j, num := range ticket {
			for ruleName, _ := range possibilities[j] {
				valid := false
				for _, rule := range rules[ruleName] {
					if rule.accepts(num) {
						valid = true
					}
				}
				if !valid {
					delete(possibilities[j], ruleName)
				}
			}
		}
	}
	return possibilities
}

func part2() int {
	ticketRules, myTicket, tickets := parseInput()
	validTickets := getValidTickets(tickets, ticketRules)
	possibilities := initPossibilities(validTickets, ticketRules)
	narrowedPossibilities := narrowPossibilities(validTickets, possibilities, ticketRules)
	for {
		allSettled := true
		for i, possibilities := range narrowedPossibilities {
			if len(possibilities) > 1 {
				allSettled = false
			} else {
				name := ""
				for k, _ := range possibilities {
					name = k
				}
				for j := 0; j < len(narrowedPossibilities); j++ {
					if j != i {
						delete(narrowedPossibilities[j], name)
					}
				}
			}
		}
		if allSettled {
			break
		}
	}
	product := 1
	for i, num := range myTicket {
		name := ""
		for k, _ := range narrowedPossibilities[i] {
			name = k
		}
		if strings.HasPrefix(name, "departure") {
			product *= num
		}
	}
	return product
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
