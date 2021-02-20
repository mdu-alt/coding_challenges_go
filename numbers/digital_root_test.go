package numbers_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/mdu-alt/coding_challenges_go/numbers"
)

func ExampleDigitalRoot() {
	// 789 -> 7+8+9 = 24 -> 2+4 = 6
	fmt.Println(numbers.DigitalRoot(789))
	// Output:
	// 6
}

func TestDigitalRoot(t *testing.T) {
	testCases := []struct {
		n    int
		want int
	}{
		// positive
		{0, 0},
		{5, 5},
		{5642, 8},
		{123456789, 9},

		// negative
		{-1, -1},
		{math.MinInt32, -1},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d", tc.n), func(t *testing.T) {
			if got := numbers.DigitalRoot(tc.n); got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}
