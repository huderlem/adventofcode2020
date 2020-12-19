// Solution for Advent of Code 2020 -- Day 19
// https://adventofcode.com/2020/day/19

package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/huderlem/adventofcode2020/util"
)

type rule interface {
	generateRegex(rules map[string]rule, repeatingRule, bothRepeatingRule string) string
}

type sequenceRule struct {
	name      string
	sequences [][]string
}

func (r sequenceRule) generateRegex(rules map[string]rule, repeatingRule string, bothRepeatingRule string) string {
	if r.name == repeatingRule {
		regex := rules[r.sequences[0][0]].generateRegex(rules, repeatingRule, bothRepeatingRule)
		return fmt.Sprintf("(%s)+", regex)
	} else if r.name == bothRepeatingRule {
		parts := []string{}
		regex1 := rules[r.sequences[0][0]].generateRegex(rules, repeatingRule, bothRepeatingRule)
		regex2 := rules[r.sequences[0][1]].generateRegex(rules, repeatingRule, bothRepeatingRule)
		// Not sure if it's possible to have two capture groups repeat the same number
		// of times in a regex, so we will just allow all possibilities up to 10 repititions.
		for i := 0; i < 10; i++ {
			c1 := ""
			c2 := ""
			for j := 0; j <= i; j++ {
				c1 += "(" + regex1 + ")"
				c2 += "(" + regex2 + ")"
			}
			parts = append(parts, fmt.Sprintf("(%s)(%s)", c1, c2))
		}
		return fmt.Sprintf("(%s)", strings.Join(parts, "|"))
	}

	parts := []string{}
	for _, sequence := range r.sequences {
		sequenceRes := []string{}
		for _, ruleName := range sequence {
			sequenceRes = append(sequenceRes, rules[ruleName].generateRegex(rules, repeatingRule, bothRepeatingRule))
		}
		parts = append(parts, strings.Join(sequenceRes, ""))
	}
	return fmt.Sprintf("(%s)", strings.Join(parts, "|"))
}

type literalRule struct {
	value byte
}

func (r literalRule) generateRegex(rules map[string]rule, repeatingRule, bothRepeatingRule string) string {
	return fmt.Sprintf("%c", r.value)
}

func parseInput() (map[string]rule, []string) {
	chunks := util.ReadFileChunks("input.txt")
	rules := map[string]rule{}
	for _, line := range chunks[0] {
		parts := strings.Split(line, ": ")
		name := parts[0]
		if strings.HasPrefix(parts[1], "\"") {
			rules[name] = literalRule{parts[1][1]}
		} else {
			sequences := [][]string{}
			curSequence := []string{}
			for _, c := range strings.Split(parts[1], " ") {
				if c == "|" {
					sequences = append(sequences, curSequence)
					curSequence = []string{}
				} else {
					curSequence = append(curSequence, c)
				}
			}
			if len(curSequence) > 0 {
				sequences = append(sequences, curSequence)
			}
			rules[name] = sequenceRule{
				name:      name,
				sequences: sequences,
			}
		}
	}

	candidates := chunks[1]
	return rules, candidates
}

func generateRulesRegex(rules map[string]rule, repeatingRule, bothRepeatingRule string) string {
	return fmt.Sprintf("^%s$", rules["0"].generateRegex(rules, repeatingRule, bothRepeatingRule))
}

func part1() int {
	rules, candidates := parseInput()
	regex := generateRulesRegex(rules, "", "")
	re, _ := regexp.Compile(regex)
	count := 0
	for _, candidate := range candidates {
		isMatch := re.MatchString(candidate)
		if isMatch {
			count++
		}
	}
	return count
}

func part2() int {
	rules, candidates := parseInput()
	regex := generateRulesRegex(rules, "8", "11")
	re, _ := regexp.Compile(regex)
	count := 0
	for _, candidate := range candidates {
		isMatch := re.MatchString(candidate)
		if isMatch {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
