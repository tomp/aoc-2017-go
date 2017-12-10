package main

import (
	"fmt"
	"github.com/tomp/aoc-2017-go/util"
    "sort"
)

const (
	INPUTFILE string = "input.txt"
)

type lineToInt func(string) (int, error)

func sumLines(lines []string, lineFn lineToInt) (int, error) {
    result := 0
    for _, line := range lines {
        lineSum, err := lineFn(line)
        if err != nil {
            return result, err
        }
        result += lineSum
    }
    return result, nil
}

func divsumLine(line string) (int, error) {
    vals, err := util.ParseInts(line)
    if err != nil {
        return 0, err
    }
    nvals := len(vals)
    sort.Ints(vals)
    for i, denom := range vals[:nvals-1] {
        for _, num := range vals[i+1:] {
            if num % denom == 0 {
                return num / denom, nil
            }
        }
    }
    return 0, fmt.Errorf("No divisible values in '%s'", line)
}

func checksumLine(line string) (int, error) {
    vals, err := util.ParseInts(line)
    if err != nil {
        return 0, err
    }
    max, min := vals[0], vals[0]
    for _, val := range vals[1:] {
        if val > max {
            max = val
        } else if val < min {
            min = val
        }
    }
    return max - min, nil
}

func checksumSheet(lines []string) (int, error) {
    return sumLines(lines, checksumLine)
}

func divsumSheet(lines []string) (int, error) {
    return sumLines(lines, divsumLine)
}

func main() {
    // Get the problem input
	lines, err := util.ReadLines(INPUTFILE)
	if err != nil {
		panic(err)
	}

	// Part 1
	checksum, err := checksumSheet(lines)
	if err != nil {
		fmt.Println("ERROR: %s", err)
	}
	fmt.Printf("1) checksum is %d\n", checksum)

	// Part 2
	divsum, err := divsumSheet(lines)
	if err != nil {
		fmt.Println("ERROR: %s", err)
	}
	fmt.Printf("2) divsum is %d\n", divsum)
}
