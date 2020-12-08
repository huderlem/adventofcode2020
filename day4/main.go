// Solution for Advent of Code 2020 -- Day 4
// https://adventofcode.com/2020/day/4

package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/huderlem/adventofcode2020/util"
)

func parsePassports() []map[string]string {
	chunks := util.ReadFileChunks("input.txt")
	re, _ := regexp.Compile(`\w+:\S+`)
	passports := []map[string]string{}
	for _, chunk := range chunks {
		curPassport := map[string]string{}
		for _, line := range chunk {
			for _, item := range re.FindAllString(line, -1) {
				parts := strings.Split(item, ":")
				curPassport[parts[0]] = parts[1]
			}
		}
		passports = append(passports, curPassport)
	}

	return passports
}

var requiredFields = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

func hasRequiredFields(passport map[string]string) bool {
	for _, field := range requiredFields {
		if _, ok := passport[field]; !ok {
			return false
		}
	}
	return true
}

func isValidRange(rawNumber string, min, max int) bool {
	num, err := strconv.Atoi(rawNumber)
	if err != nil {
		return false
	}
	return num >= min && num <= max
}

func isValidHeight(rawHeight string) bool {
	l := len(rawHeight)
	num := rawHeight[:l-2]
	units := rawHeight[l-2:]
	if units == "cm" {
		return isValidRange(num, 150, 193)
	} else if units == "in" {
		return isValidRange(num, 59, 76)
	}
	return false
}

func isValidHairColor(hairColor string) bool {
	if hairColor[0] != '#' {
		return false
	}
	re, _ := regexp.Compile(`^[a-z0-9]{6}$`)
	return re.MatchString(hairColor[1:])
}

func isValidEyeColor(eyeColor string) bool {
	return eyeColor == "amb" || eyeColor == "blu" ||
		eyeColor == "brn" || eyeColor == "gry" ||
		eyeColor == "grn" || eyeColor == "hzl" ||
		eyeColor == "oth"
}

func isValidPassportId(id string) bool {
	re, _ := regexp.Compile(`^[0-9]{9}$`)
	return re.MatchString(id)
}

func part1() int {
	passports := parsePassports()
	numValid := 0
	for _, passport := range passports {
		if hasRequiredFields(passport) {
			numValid++
		}
	}
	return numValid
}

func part2() int {
	passports := parsePassports()
	numValid := 0
	for _, passport := range passports {
		if hasRequiredFields(passport) &&
			isValidRange(passport["byr"], 1920, 2002) &&
			isValidRange(passport["iyr"], 2010, 2020) &&
			isValidRange(passport["eyr"], 2020, 2030) &&
			isValidHeight(passport["hgt"]) &&
			isValidHairColor(passport["hcl"]) &&
			isValidEyeColor(passport["ecl"]) &&
			isValidPassportId(passport["pid"]) {
			numValid++
		}
	}
	return numValid
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
