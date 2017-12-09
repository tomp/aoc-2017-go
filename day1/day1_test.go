package main

import (
	"testing"
)

func TestSumRepeats1(t *testing.T) {
	cases := [...]struct {
		text     string
		expected int
	}{
		{"1122", 3},
		{"1111", 4},
		{"1234", 0},
		{"91212129", 9},
	}

	for _, item := range cases {
		result, err := sumRepeats1(item.text)
		if err != nil {
			t.Errorf("error for '%s': %s", item.text, err)
		}
		if result != item.expected {
			t.Errorf("sum for '%s' is %d (expected %d)", item.text, result, item.expected)
		}
	}
}

func TestSumRepeats(t *testing.T) {
	cases := [...]struct {
		text     string
		expected int
	}{
		{"1212", 6},
		{"1221", 0},
		{"123425", 4},
		{"123123", 12},
		{"12131415", 4},
	}

	for _, item := range cases {
		result, err := sumRepeats(item.text, len(item.text) / 2)
		if err != nil {
			t.Errorf("error for '%s': %s", item.text, err)
		}
		if result != item.expected {
			t.Errorf("sum for '%s' is %d (expected %d)", item.text, result, item.expected)
		}
	}
}
