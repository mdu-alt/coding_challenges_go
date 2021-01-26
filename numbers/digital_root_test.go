package numbers_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/mdu-alt/coding_challenges_go/numbers"
)

func ExampleDigitalRoot() {
	// 789 -> 7+8+9 == 24 -> 2+4 == 6
	fmt.Println(numbers.DigitalRoot(789))
	// Output: 6
}

func TestDigitalRoot(t *testing.T) {
	testMatrix := map[string][]struct {
		n    int
		want int
	}{
		"positive": {
			{0, 0},
			{5, 5},
			{5642, 8},
			{123456789, 9},
		},
		"negative": {
			{-1, -1},
			{math.MinInt32, -1},
		},
	}

	for title, testCases := range testMatrix {
		for _, test := range testCases {
			name := fmt.Sprintf("%s{%d|%d}", title, test.n, test.want)

			t.Run(name, func(t *testing.T) {
				if got := numbers.DigitalRoot(test.n); got != test.want {
					t.Errorf("expected: %d, got: %d", test.want, got)
				}
			})
		}
	}
}
