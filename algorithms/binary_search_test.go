package algorithms

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	testCases := map[string][]struct {
		s        []int
		e        int
		expected int
	}{
		"valid": {
			{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5, 4},
			{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 7, 6},
			{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 1, 0},
			{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9, 8},
		},
		"invalid": {
			{[]int{1, 2, 3, 4, 4, 6, 7, 8, 9}, 5, -1},
			{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 0, -1},
			{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 10, -1},
		},
	}

	for name, tC := range testCases {
		for _, test := range tC {
			testname := fmt.Sprintf("%v--{%v,%v,%v}", name, test.s, test.e, test.expected)

			t.Run(testname, func(t *testing.T) {
				if got := binarySearch(test.s, test.e); got != test.expected {
					t.Errorf("expected: %v, got: %v", test.expected, got)
				}
			})
		}
	}
}
