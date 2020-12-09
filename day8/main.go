// Solution for Advent of Code 2020 -- Day 8
// https://adventofcode.com/2020/day/8

package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/huderlem/adventofcode2020/util"
)

type instruction struct {
	op    string
	param int
}

func parseProgram() []instruction {
	lines := util.ReadFileLines("input.txt")
	program := []instruction{}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		param, _ := strconv.Atoi(parts[1][1:])
		if parts[1][0] == '-' {
			param = -param
		}
		program = append(program, instruction{
			op:    parts[0],
			param: param,
		})
	}
	return program
}

func runProgram(program []instruction) (bool, int) {
	acc := 0
	pc := 0
	visited := map[int]bool{}
	for {
		if pc == len(program) {
			// Program terminated successfully.
			return true, acc
		}
		if _, ok := visited[pc]; ok {
			// Encountered an infinite loop.
			return false, acc
		}
		visited[pc] = true
		switch program[pc].op {
		case "acc":
			acc += program[pc].param
			pc++
		case "jmp":
			pc += program[pc].param
		case "nop":
			pc++
		}
	}
}

func part1() int {
	program := parseProgram()
	terminates, acc := runProgram(program)
	if terminates {
		// Program is supposed to result in an infinite loop.
		return -1
	}
	return acc
}

func part2() int {
	program := parseProgram()
	// Brute-force and try flipping each "jmp" or "nop" instruction in the
	// entire program, one at a time.
	for i, instr := range program {
		if instr.op == "jmp" || instr.op == "nop" {
			modifiedProgram := make([]instruction, len(program))
			copy(modifiedProgram, program)
			if instr.op == "jmp" {
				modifiedProgram[i].op = "nop"
			} else {
				modifiedProgram[i].op = "jmp"
			}
			terminates, acc := runProgram(modifiedProgram)
			if terminates {
				return acc
			}
		}
	}
	// None of the flipped instructions result in a terminating program.
	return -1
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
