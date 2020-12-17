// Solution for Advent of Code 2020 -- Day 17
// https://adventofcode.com/2020/day/17

package main

import (
	"fmt"
	"math"

	"github.com/huderlem/adventofcode2020/util"
)

type point struct {
	x, y, z, w int
}

func parseInput() map[point]struct{} {
	lines := util.ReadFileLines("input.txt")
	cells := map[point]struct{}{}
	for y, line := range lines {
		for x, cell := range line {
			if cell == '#' {
				p := point{x, y, 0, 0}
				cells[p] = struct{}{}
			}
		}
	}
	return cells
}

func getBounds(cells map[point]struct{}) (point, point) {
	min := point{math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}
	max := point{math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64}
	for cell, _ := range cells {
		if cell.x < min.x {
			min.x = cell.x
		}
		if cell.y < min.y {
			min.y = cell.y
		}
		if cell.z < min.z {
			min.z = cell.z
		}
		if cell.w < min.w {
			min.w = cell.w
		}
		if cell.x > max.x {
			max.x = cell.x
		}
		if cell.y > max.y {
			max.y = cell.y
		}
		if cell.z > max.z {
			max.z = cell.z
		}
		if cell.w > max.w {
			max.w = cell.w
		}
	}
	min.x--
	min.y--
	min.z--
	min.w--
	max.x++
	max.y++
	max.z++
	max.w++
	return min, max
}

func countActiveNeighbors(p point, cells map[point]struct{}) int {
	count := 0
	for x := p.x - 1; x <= p.x+1; x++ {
		for y := p.y - 1; y <= p.y+1; y++ {
			for z := p.z - 1; z <= p.z+1; z++ {
				for w := p.w - 1; w <= p.w+1; w++ {
					if x == p.x && y == p.y && z == p.z && w == p.w {
						continue
					}
					neighbor := point{x, y, z, w}
					if _, ok := cells[neighbor]; ok {
						count++
					}
				}
			}
		}
	}
	return count
}

func applyRules3d(cells map[point]struct{}) map[point]struct{} {
	newCells := map[point]struct{}{}
	for k, _ := range cells {
		newCells[k] = struct{}{}
	}
	min, max := getBounds(cells)
	for x := min.x; x <= max.x; x++ {
		for y := min.y; y <= max.y; y++ {
			for z := min.z; z <= max.z; z++ {
				p := point{x, y, z, 0}
				activeNeighbors := countActiveNeighbors(p, cells)
				if _, ok := cells[p]; ok {
					if activeNeighbors != 2 && activeNeighbors != 3 {
						delete(newCells, p)
					}
				} else {
					if activeNeighbors == 3 {
						newCells[p] = struct{}{}
					}
				}
			}
		}
	}

	return newCells
}

func applyRules4d(cells map[point]struct{}) map[point]struct{} {
	newCells := map[point]struct{}{}
	for k, _ := range cells {
		newCells[k] = struct{}{}
	}
	min, max := getBounds(cells)
	for x := min.x; x <= max.x; x++ {
		for y := min.y; y <= max.y; y++ {
			for z := min.z; z <= max.z; z++ {
				for w := min.w; w <= max.w; w++ {
					p := point{x, y, z, w}
					activeNeighbors := countActiveNeighbors(p, cells)
					if _, ok := cells[p]; ok {
						if activeNeighbors != 2 && activeNeighbors != 3 {
							delete(newCells, p)
						}
					} else {
						if activeNeighbors == 3 {
							newCells[p] = struct{}{}
						}
					}
				}
			}
		}
	}

	return newCells
}

func part1() int {
	cells := parseInput()
	for i := 0; i < 6; i++ {
		cells = applyRules3d(cells)
	}
	return len(cells)
}

func part2() int {
	cells := parseInput()
	for i := 0; i < 6; i++ {
		cells = applyRules4d(cells)
	}
	return len(cells)
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
