// Solution for Advent of Code 2020 -- Day 18
// https://adventofcode.com/2020/day/18

package main

import (
	"fmt"
	"strconv"

	"github.com/huderlem/adventofcode2020/util"
)

type token struct {
	kind  string
	value int
}

func tokenizeInput() [][]token {
	lines := util.ReadFileLines("input.txt")
	tokenizedLines := [][]token{}
	for _, line := range lines {
		tokenizedLine := []token{}
		for _, c := range line {
			if c == ' ' {
				continue
			}
			t := token{}
			switch c {
			case '(':
				t.kind = "("
			case ')':
				t.kind = ")"
			case '+':
				t.kind = "+"
			case '*':
				t.kind = "*"
			default:
				t.kind = "INT"
				t.value, _ = strconv.Atoi(string(c))
			}
			tokenizedLine = append(tokenizedLine, t)
		}
		tokenizedLines = append(tokenizedLines, tokenizedLine)
	}
	return tokenizedLines
}

// Shunting-yard algorithm for generating postfix notation.
// https://en.wikipedia.org/wiki/Shunting-yard_algorithm
//
// I happen to have used this in the past for porymap's constant
// evaluation--how convenient!
// https://github.com/huderlem/porymap/blob/master/src/core/parseutil.cpp#L156
func generatePostfixTokens(tokens []token, precedence map[string]int) []token {
	postfix := []token{}
	opStack := []token{}
	for _, t := range tokens {
		switch t.kind {
		case "INT":
			postfix = append(postfix, t)
		case "(":
			opStack = append(opStack, t)
		case ")":
			for len(opStack) > 0 && opStack[len(opStack)-1].kind != "(" {
				postfix = append(postfix, opStack[len(opStack)-1])
				opStack = opStack[:len(opStack)-1]
			}
			if len(opStack) > 0 {
				// pop the left parenthesis token
				opStack = opStack[:len(opStack)-1]
			} else {
				panic("Mismatched parentheses detected in expression!")
			}
		default:
			// token is an operator
			for len(opStack) > 0 && precedence[opStack[len(opStack)-1].kind] <= precedence[t.kind] && opStack[len(opStack)-1].kind != "(" {
				postfix = append(postfix, opStack[len(opStack)-1])
				opStack = opStack[:len(opStack)-1]
			}
			opStack = append(opStack, t)
		}
	}

	for len(opStack) > 0 {
		if opStack[len(opStack)-1].kind == "(" || opStack[len(opStack)-1].kind == ")" {
			panic("Mismatched parentheses detected in epxression!")
		}
		postfix = append(postfix, opStack[len(opStack)-1])
		opStack = opStack[:len(opStack)-1]
	}

	return postfix
}

// Evaluate postfix expression.
// https://en.wikipedia.org/wiki/Reverse_Polish_notation#Postfix_evaluation_algorithm
func evaluatePostfix(postfix []token) int {
	stack := []token{}
	for _, t := range postfix {
		if t.kind == "*" || t.kind == "+" && len(stack) > 1 {
			op2 := stack[len(stack)-1].value
			op1 := stack[len(stack)-2].value
			stack = stack[:len(stack)-2]
			result := 0
			switch t.kind {
			case "*":
				result = op1 * op2
			case "+":
				result = op1 + op2
			}
			stack = append(stack, token{kind: "INT", value: result})
		} else {
			stack = append(stack, t)
		}
	}

	return stack[0].value
}
func part1() int {
	tokenizedExpressions := tokenizeInput()
	precedence := map[string]int{
		"+": 1,
		"*": 1,
	}
	sum := 0
	for _, tokens := range tokenizedExpressions {
		postfix := generatePostfixTokens(tokens, precedence)
		sum += evaluatePostfix(postfix)
	}
	return sum
}

func part2() int {
	tokenizedExpressions := tokenizeInput()
	precedence := map[string]int{
		"+": 1,
		"*": 2,
	}
	sum := 0
	for _, tokens := range tokenizedExpressions {
		postfix := generatePostfixTokens(tokens, precedence)
		sum += evaluatePostfix(postfix)
	}
	return sum
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
