package main

import (
	"fmt"

	"github.com/ArtKelly/AoC2020Golang/util"
)

func main() {
	lines1, err := util.ReadLines("Challenges/Day1/1.txt")
	util.Check(err)

	lines, err := util.StringsArrayToInts(lines1)
	util.Check(err)

	// PART 1
	result, err := twoSum(lines, 2020)
	util.Check(err)
	fmt.Printf("Part 1 Solution: %d\n", result[0]*result[1])

	// PART 2
	result, err = threeSum(lines, 2020)
	util.Check(err)
	fmt.Printf("Part 2 Solution: %d\n", result[0]*result[1]*result[2])
}

func twoSum(arr []int, sum int) ([]int, error) {
	var result []int
	var solution bool = false

	for i := 0; i < len(arr)-1; i++ {
		for j := i; j < len(arr)-1; j++ {
			if arr[i]+arr[j] == 2020 {
				result = append(result, arr[i], arr[j])
				solution = true
				break
			}
		}
	}

	if solution {
		return result, nil
	}

	return []int{0, 0}, fmt.Errorf("no pair found to make sum of %d", sum)
}

func threeSum(arr []int, sum int) ([]int, error) {
	var result []int
	var solution bool = false

	for i := 0; i < len(arr)-1; i++ {
		for j := i; j < len(arr)-1; j++ {
			for k := j; k < len(arr)-1; k++ {
				if arr[i]+arr[j]+arr[k] == 2020 {
					result = append(result, arr[i], arr[j], arr[k])
					solution = true
					break
				}
			}
		}
	}

	if solution {
		return result, nil
	}
	return []int{0, 0, 0}, fmt.Errorf("no triple found to make sum of %d", sum)
}
