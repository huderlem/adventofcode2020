// Solution for Advent of Code 2020 -- Day 3
// https://adventofcode.com/2020/day/3

package main

import (
	"fmt"

	"github.com/huderlem/adventofcode2020/util"
)

func parseTreesGrid() [][]bool {
	lines := util.ReadFileLines("input.txt")
	width := len(lines[0])
	grid := make([][]bool, width)
	for _, line := range lines {
		for x, cell := range line {
			grid[x] = append(grid[x], byte(cell) == '#')
		}
	}
	return grid
}

func countHitTrees(grid [][]bool, deltaX, deltaY int) int {
	width := len(grid)
	height := len(grid[0])
	x := 0
	y := 0
	numTreesHit := 0
	for y < height {
		if grid[x][y] {
			numTreesHit++
		}
		x = (x + deltaX) % width
		y += deltaY
	}

	return numTreesHit
}

func part1() int {
	grid := parseTreesGrid()
	return countHitTrees(grid, 3, 1)
}

func part2() int {
	grid := parseTreesGrid()
	return countHitTrees(grid, 1, 1) * countHitTrees(grid, 3, 1) * countHitTrees(grid, 5, 1) * countHitTrees(grid, 7, 1) * countHitTrees(grid, 1, 2)
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
