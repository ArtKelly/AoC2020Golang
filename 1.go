package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	lines1, err := readLines("1.txt")
	check(err)
	var lines = []int{}

	for _, i := range lines1 {
		j, err := strconv.Atoi(i)
		check(err)
		lines = append(lines, j)
	}

	for i := 0; i < len(lines)-1; i++ {
		for j := i; j < len(lines)-1; j++ {
			if lines[i]+lines[j] == 2020 {
				fmt.Println(lines[i] * lines[j])
			}
		}
	}
}
