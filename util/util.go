package util

import (
	"bufio"
	"os"
)

// ReadLines returns the contents of the given file as a slice
// of lines.
func ReadLines(filename string) (lines []string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}
