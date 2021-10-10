package main

import (
	"fmt"

	"github.com/ArtKelly/AoC2020Golang/util"
)

func main() {
	lines, err := util.ReadLines("Challenges/Day3/input.txt")
	util.Check(err)

	trees := countTrees(lines, 3, 1)
	fmt.Printf("Part 1 trees: %d\n", trees)

	part2 := 1 * countTrees(lines, 1, 1) * countTrees(lines, 3, 1) * countTrees(lines, 5, 1) * countTrees(lines, 7, 1) * countTrees(lines, 1, 2)
	fmt.Printf("Part 2 trees: %d\n", part2)
}

func countTrees(lines []string, dx int, dy int) int {
	var trees int = 0
	for x, y := 0, 0; y < len(lines); x, y = x+dx, y+dy {
		if string(lines[y][x%len(lines[y])]) == "#" {
			trees++
		}
	}
	return trees
}
