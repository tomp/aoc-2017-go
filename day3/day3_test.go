package main

import (
	"testing"
)

func TestDistanceToOrigin(t *testing.T) {
	cases := [...]struct {
		start    int
		expected int
	}{
		{1, 0},
		{12, 3},
		{23, 2},
		{1024, 31},
	}

	for _, item := range cases {
		result := distanceToOrigin(item.start)
		if result != item.expected {
			t.Errorf("for start=%d, result is %d (expected %d)",
				item.start, result, item.expected)
		}
	}
}
