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
	fmt.Printf("Part 1: %d", count)
}

func getBounds(str string) (int, int) {
	s := strings.Split(str, "-")

	lowerBound, err := strconv.Atoi(s[0])
	util.Check(err)

	upperBound, err := strconv.Atoi(s[1])
	util.Check(err)

	return lowerBound, upperBound
}
