package numbers

import (
	"fmt"
	"math"
	"testing"
)

func TestLargestBinaryGap(t *testing.T) {
	type testCase struct {
		n        int
		expected int
	}

	testMatrix := map[string][]testCase{
		"any_gaps": {
			{5, 1},          //          5 = 0b101
			{9, 2},          //          9 = 0b1001
			{17, 3},         //         17 = 0b10001
			{20, 1},         //         20 = 0b10100
			{2546, 2},       //       2546 = 0b100111110010
			{1376796946, 5}, // 1376796946 = 0b1010010000100000100000100010010
		},
		"no_gaps": {
			{1, 0},  //  1 = 0b1
			{3, 0},  //  3 = 0b11
			{7, 0},  //  7 = 0b111
			{14, 0}, // 14 = 0b1110
			{24, 0}, // 24 = 0b11000
			{32, 0}, // 32 = 0b100000
		},
		"extremes": {
			{0, 0},             //   0 = 0b0
			{math.MaxInt32, 0}, // Max ~ 0b1111111111...
		},
		"invalid": {
			{-1, -1},
			{math.MinInt32, -1},
		},
	}

	for name, tests := range testMatrix {
		for _, test := range tests {
			testname := fmt.Sprintf("%v--{%v,%v}", name, test.n, test.expected)

			t.Run(testname, func(t *testing.T) {
				if got := largestBinaryGap(test.n); got != test.expected {
					t.Errorf("expected: %v, got: %v", test.expected, got)
				}
			})
		}
	}
}
