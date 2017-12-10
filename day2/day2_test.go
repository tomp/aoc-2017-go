package main

import (
	"testing"
)

func TestChecksumSheet(t *testing.T) {
    lines := []string{"5 1 9 5",
                      "7 5 3",
                      "2 4 6 8"}
    expected := 18

    result, err := checksumSheet(lines)
    if err != nil {
        t.Errorf("error: %s", err)
    }
    if result != expected {
        t.Errorf("checksum for sheet is %d (expected %d)", result, expected)
    }
}

func TestChecksumLine(t *testing.T) {
	cases := [...]struct {
		text     string
		expected int
	}{
		{"5 1 9 5", 8},
		{"7 5 3", 4},
		{"2 4 6 8", 6},
	}

	for _, item := range cases {
		result, err := checksumLine(item.text)
		if err != nil {
			t.Errorf("error for '%s': %s", item.text, err)
		}
		if result != item.expected {
			t.Errorf("checksum for '%s' is %d (expected %d)", item.text, result, item.expected)
		}
	}
}

/*
func TestDivsumSheet(t *testing.T) {
    lines := []string{"5 9 2 8",
                      "9 4 7 3",
                      "3 8 6 5"}
    expected := 9

    result, err := divsumSheet(lines)
    if err != nil {
        t.Errorf("error: %s", err)
    }
    if result != expected {
        t.Errorf("divsum for sheet is %d (expected %d)", result, expected)
    }
}
*/

func TestDivsumLine(t *testing.T) {
	cases := [...]struct {
		text     string
		expected int
	}{
		{"5 9 2 8", 4},
		{"9 4 7 3", 3},
		{"3 8 6 5", 2},
	}

	for _, item := range cases {
		result, err := divsumLine(item.text)
		if err != nil {
			t.Errorf("error for '%s': %s", item.text, err)
		}
		if result != item.expected {
			t.Errorf("divsum for '%s' is %d (expected %d)", item.text, result, item.expected)
		}
	}
}
