package day01

import (
	"testing"
)

const (
	Example     = "example.txt"
	PuzzleInput = "puzzle_input.txt"
)

func TestPart1(t *testing.T) {
	testCases := []struct {
		filename string
		want     int
	}{
		{Example, 7},
		{PuzzleInput, 1766},
	}

	for _, tc := range testCases {
		t.Run(tc.filename, func(t *testing.T) {
			if got := part1(tc.filename); got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	testCases := []struct {
		filename string
		want     int
	}{
		{Example, 5},
		{PuzzleInput, 1797},
	}

	for _, tc := range testCases {
		t.Run(tc.filename, func(t *testing.T) {
			if got := part2(tc.filename); got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
