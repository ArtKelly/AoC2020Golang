package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ArtKelly/AoC2020Golang/util"
)

func main() {
	lines, err := util.ReadLines("Challenges/Day2/input.txt")
	util.Check(err)

	count := partOne(lines)
	fmt.Printf("Part 1: %d\n", count)

	count = partTwo(lines)
	fmt.Printf("Part 2: %d\n", count)
}

func getBounds(str string) (int, int) {
	s := strings.Split(str, "-")

	lowerBound, err := strconv.Atoi(s[0])
	util.Check(err)

	upperBound, err := strconv.Atoi(s[1])
	util.Check(err)

	return lowerBound, upperBound
}

func partOne(lines []string) int {
	var currentPassword []string
	var count int = 0

	for i := 0; i < len(lines); i++ {
		currentPassword = strings.Fields(lines[i])
		rule, character, password := currentPassword[0], strings.TrimSuffix(currentPassword[1], ":"), currentPassword[2]
		lowerBound, upperBound := getBounds(rule)
		if lowerBound <= strings.Count(password, character) && strings.Count(password, character) <= upperBound {
			count++
		}
	}
	return count
}

func partTwo(lines []string) int {
	var currentPassword []string
	var count int = 0
	var a, b bool = false, false

	for i := 0; i < len(lines); i++ {
		currentPassword = strings.Fields(lines[i])
		rule, character, password := currentPassword[0], strings.TrimSuffix(currentPassword[1], ":"), currentPassword[2]
		lowerBound, upperBound := getBounds(rule)
		a, b = string(password[lowerBound-1]) == character, string(password[upperBound-1]) == character
		if a != b {
			count++
		}
	}
	return count
}
