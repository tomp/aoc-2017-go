package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
    lines := []string{"5 1 9 5",
                      "7 5 3",
                      "2 4 6 8"}
    expected := 18

    result, err := solve(lines)
    if err != nil {
        t.Errorf("error: %s", err)
    }
    if result != expected {
        t.Errorf("result is %d (expected %d)", result, expected)
    }
}

