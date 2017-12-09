package main

import (
	"fmt"
	"github.com/tomp/aoc-2017-go/util"
	"strconv"
)

const (
	INPUTFILE string = "input.txt"
)

// parseDigits takes a string consisting of integer digits
// and returns a slice of the integer value for each digit.
func parseDigits(text string) (digits []int, err error) {
	digits = make([]int, len(text))
	for i, ch := range text {
		val, err := strconv.Atoi(string(ch))
		if err != nil {
			return digits, err
		}
		digits[i] = val
	}
	return digits, nil
}

func sumRepeats1(text string) (total int, err error) {
    return sumRepeats(text, 1)
}

func sumRepeats(text string, gap int) (total int, err error) {
	n := len(text)
	total = 0
	digits, err := parseDigits(text)
	if err != nil {
		return total, err
	}
	for i, digit := range digits {
		j := (i + gap) % n
		if digit == digits[j] {
			total = total + digit
		}
	}
	return total, nil
}

func main() {
    // Get the problem input
	lines, err := util.ReadLines(INPUTFILE)
	if err != nil {
		panic(err)
	}
    text := lines[0]

	// Part 1
	count, err := sumRepeats1(text)
	if err != nil {
		fmt.Println("ERROR: %s", err)
	}
	fmt.Printf("1) repeats sum to %d\n", count)

	// Part 2
	count, err = sumRepeats(lines[0], len(text)/2)
	if err != nil {
		fmt.Println("ERROR: %s", err)
	}
	fmt.Printf("2) repeats sum to %d\n", count)
}
