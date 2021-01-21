package numbers

import (
	"fmt"
	"math"
	"testing"
)

func TestDigitalRoot(t *testing.T) {
	type testCase struct {
		n        int
		expected int
	}

	testMatrix := map[string][]testCase{
		"valid": {
			{0, 0},
			{5, 5},
			{39, 3},
			{871, 7},
			{5642, 8},
			{123456789, 9},
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
				if got := digitalRoot(test.n); got != test.expected {
					t.Errorf("expected: %v, got: %v", test.expected, got)
				}
			})
		}
	}
}
