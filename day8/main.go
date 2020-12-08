// Solution for Advent of Code 2020 -- Day 7
// https://adventofcode.com/2020/day/7

package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/huderlem/adventofcode2020/util"
)

type instr struct {
	op    string
	param int
}

func parseInput() []instr {
	lines := util.ReadFileLines("input.txt")
	instructions := []instr{}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		i := instr{
			op: parts[0],
		}
		v, _ := strconv.Atoi(parts[1][1:])
		if parts[1][0] == '+' {
			i.param = v
		} else {
			i.param = -v
		}
		instructions = append(instructions, i)
	}
	return instructions
}

func part1() int {
	instructions := parseInput()
	flipped := []int{}
	for i, inst := range instructions {
		if inst.op == "jmp" || inst.op == "nop" {
			flipped = append(flipped, i)
		}
	}
	for _, i := range flipped {
		acc := 0
		pc := 0
		visited := map[int]bool{}
		for {
			if pc == len(instructions) {
				return acc
			}
			if _, ok := visited[pc]; ok {
				break
			}
			visited[pc] = true
			op := instructions[pc].op
			if pc == i {
				if op == "acc" {
					op = "jmp"
				} else if op == "jmp" {
					op = "nop"
				}
			}
			switch op {
			case "acc":
				acc += instructions[pc].param
				pc++
			case "jmp":
				pc += instructions[pc].param
			case "nop":
				pc++
			}
		}
	}
	return -1
}

func part2() int {
	instructions := parseInput()
	flipped := []int{}
	for i, inst := range instructions {
		if inst.op == "jmp" || inst.op == "nop" {
			flipped = append(flipped, i)
		}
	}
	for _, i := range flipped {
		acc := 0
		pc := 0
		visited := map[int]bool{}
		for {
			if pc == len(instructions) {
				return acc
			}
			if _, ok := visited[pc]; ok {
				break
			}
			visited[pc] = true
			op := instructions[pc].op
			if pc == i {
				if op == "acc" {
					op = "jmp"
				} else if op == "jmp" {
					op = "nop"
				}
			}
			switch op {
			case "acc":
				acc += instructions[pc].param
				pc++
			case "jmp":
				pc += instructions[pc].param
			case "nop":
				pc++
			}
		}
	}
	return -1
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
