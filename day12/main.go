// Solution for Advent of Code 2020 -- Day 12
// https://adventofcode.com/2020/day/12

package main

import (
	"fmt"
	"strconv"

	"github.com/huderlem/adventofcode2020/util"
)

type direction int

const (
	north direction = iota
	east
	south
	west
)

var directions = []direction{north, east, south, west}

func (d direction) rotate(degrees int) direction {
	steps := (degrees / 90)
	newOffset := (int(d) + steps) % len(directions)
	if newOffset < 0 {
		newOffset += len(directions)
	}
	return directions[newOffset]
}

type command struct {
	action string
	value  int
}

type point struct {
	x, y int
}

func (p *point) rotateRight(amount int) {
	for steps := (amount / 90) % 4; steps > 0; steps-- {
		p.x, p.y = p.y, -p.x
	}
}

func (p *point) rotateLeft(amount int) {
	for steps := (amount / 90) % 4; steps > 0; steps-- {
		p.x, p.y = -p.y, p.x
	}
}

func (p *point) moveNorth(amount int) {
	p.y += amount
}

func (p *point) moveEast(amount int) {
	p.x += amount
}

func (p *point) moveSouth(amount int) {
	p.y -= amount
}

func (p *point) moveWest(amount int) {
	p.x -= amount
}

type boat struct {
	pos      point
	facing   direction
	waypoint point
}

func (b *boat) turnLeft(degrees int) {
	b.facing = b.facing.rotate(-degrees)
}

func (b *boat) turnRight(degrees int) {
	b.facing = b.facing.rotate(degrees)
}

func (b *boat) moveTowardWaypoint(amount int) {
	b.pos.x += b.waypoint.x * amount
	b.pos.y += b.waypoint.y * amount
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func parseCommands() []command {
	lines := util.ReadFileLines("input.txt")
	commands := make([]command, len(lines))
	for i, line := range lines {
		action := line[:1]
		value, _ := strconv.Atoi(line[1:])
		commands[i] = command{action, value}
	}
	return commands
}

func part1() int {
	commands := parseCommands()
	boat := boat{pos: point{0, 0}, facing: east}
	for _, command := range commands {
		switch command.action {
		case "N":
			boat.pos.moveNorth(command.value)
		case "E":
			boat.pos.moveEast(command.value)
		case "S":
			boat.pos.moveSouth(command.value)
		case "W":
			boat.pos.moveWest(command.value)
		case "L":
			boat.turnLeft(command.value)
		case "R":
			boat.turnRight(command.value)
		case "F":
			switch boat.facing {
			case north:
				boat.pos.moveNorth(command.value)
			case east:
				boat.pos.moveEast(command.value)
			case south:
				boat.pos.moveSouth(command.value)
			case west:
				boat.pos.moveWest(command.value)
			}
		}
	}

	return abs(boat.pos.x) + abs(boat.pos.y)
}

func part2() int {
	commands := parseCommands()
	boat := boat{pos: point{0, 0}, waypoint: point{10, 1}}
	for _, command := range commands {
		switch command.action {
		case "N":
			boat.waypoint.moveNorth(command.value)
		case "E":
			boat.waypoint.moveEast(command.value)
		case "S":
			boat.waypoint.moveSouth(command.value)
		case "W":
			boat.waypoint.moveWest(command.value)
		case "L":
			boat.waypoint.rotateLeft(command.value)
		case "R":
			boat.waypoint.rotateRight(command.value)
		case "F":
			boat.moveTowardWaypoint(command.value)
		}
	}

	return abs(boat.pos.x) + abs(boat.pos.y)
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
