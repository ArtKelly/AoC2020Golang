package main

import (
	"fmt"
	"strconv"

	"github.com/ArtKelly/AoC2020Golang/util"
)

func main() {
	lines1, err := util.ReadLines("Challenges/Day1/1.txt")
	util.Check(err)
	var lines = []int{}

	for _, i := range lines1 {
		j, err := strconv.Atoi(i)
		util.Check(err)
		lines = append(lines, j)
	}

	for i := 0; i < len(lines)-1; i++ {
		for j := i; j < len(lines)-1; j++ {
			if lines[i]+lines[j] == 2020 {
				fmt.Println(lines[i] * lines[j])
				break
			}
		}
	}
}
