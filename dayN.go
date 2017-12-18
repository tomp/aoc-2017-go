package main

import (
	"fmt"
	"github.com/tomp/aoc-2017-go/util"
)

const (
	INPUTFILE string = "input.txt"
)

func solve(lines []string) int {
    return 0
}

func solve2(lines []string) int {
    return 0
}

func main() {
    // Get the problem input
	lines, err := util.ReadLines(INPUTFILE)
	if err != nil {
		panic(err)
	}

	// Part 1
	result, err := solve(lines)
	if err != nil {
		fmt.Println("ERROR: %s", err)
	}
	fmt.Printf("1) result is %d\n", result)

	// Part 2
	result, err := solve2(lines)
	if err != nil {
		fmt.Println("ERROR: %s", err)
	}
	fmt.Printf("2) result is %d\n", result)
}
