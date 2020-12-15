// Solution for Advent of Code 2020 -- Day 13
// https://adventofcode.com/2020/day/13

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/huderlem/adventofcode2020/util"
)

func setBit(value uint64, bitIndex int) uint64 {
	return value | (1 << bitIndex)
}

func resetBit(value uint64, bitIndex int) uint64 {
	return value & (math.MaxUint64 ^ (1 << bitIndex))
}

func applyValueMask(value uint64, mask string) uint64 {
	for i, c := range mask {
		if c == 'X' {
			continue
		}
		bitIndex := len(mask) - i - 1
		if c == '0' {
			value = resetBit(value, bitIndex)
		} else if c == '1' {
			value = setBit(value, bitIndex)
		}
	}
	return value
}

func applyAddressMask(address uint64, mask string) []uint64 {
	for i, c := range mask {
		if c == '1' {
			bitIndex := len(mask) - i - 1
			address = setBit(address, bitIndex)
		}
	}
	floatingBitIndexes := []int{}
	for i, c := range mask {
		if c == 'X' {
			bitIndex := len(mask) - i - 1
			floatingBitIndexes = append(floatingBitIndexes, bitIndex)
		}
	}
	numAddresses := int(math.Pow(2, float64(len(floatingBitIndexes))))
	addresses := make([]uint64, numAddresses)
	for i := 0; i < numAddresses; i++ {
		maskedAddress := address
		for j, bitIndex := range floatingBitIndexes {
			if i&(1<<j) == 0 {
				maskedAddress = resetBit(maskedAddress, bitIndex)
			} else {
				maskedAddress = setBit(maskedAddress, bitIndex)
			}
		}
		addresses[i] = maskedAddress
	}
	return addresses
}

func part1() uint64 {
	lines := util.ReadFileLines("input.txt")
	mask := ""
	memory := map[uint64]uint64{}
	for _, line := range lines {
		if strings.HasPrefix(line, "mask") {
			mask = line[7:]
		} else if strings.HasPrefix(line, "mem") {
			address, _ := strconv.ParseUint(line[4:strings.Index(line, "]")], 10, 64)
			value, _ := strconv.ParseUint(line[strings.Index(line, "=")+2:], 10, 64)
			maskedValue := applyValueMask(value, mask)
			memory[address] = maskedValue
		}
	}
	var sum uint64
	for _, v := range memory {
		sum += v
	}

	return sum
}

func part2() uint64 {
	lines := util.ReadFileLines("input.txt")
	mask := ""
	memory := map[uint64]uint64{}
	for _, line := range lines {
		if strings.HasPrefix(line, "mask") {
			mask = line[7:]
		} else if strings.HasPrefix(line, "mem") {
			address, _ := strconv.ParseUint(line[4:strings.Index(line, "]")], 10, 64)
			value, _ := strconv.ParseUint(line[strings.Index(line, "=")+2:], 10, 64)
			addresses := applyAddressMask(address, mask)
			for _, floatingAddress := range addresses {
				memory[floatingAddress] = value
			}
		}
	}
	var sum uint64
	for _, v := range memory {
		sum += v
	}

	return sum
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
