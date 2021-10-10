package main

import (
	"fmt"

	"github.com/ArtKelly/AoC2020Golang/util"
)

func main() {
	lines, err := util.ReadLines("Challenges/Day3/input.txt")
	util.Check(err)
	var trees int = 0
	for x, y := 0, 0; y < len(lines); x, y = x+3, y+1 {
		if string(lines[y][x%len(lines[y])]) == "#" {
			trees++
		}
	}
	fmt.Printf("Trees: %d\n", trees)
}
