package numbers

import (
	"fmt"
	"testing"
)

func ExampleDigitalRoot() {
	// 789 -> 7+8+9 = 24 -> 2+4 = 6
	fmt.Println(DigitalRoot(789))
	// Output:
	// 6
}

func TestDigitalRoot(t *testing.T) {
	testCases := []struct {
		n    int
		want int
	}{
		{0, 0},
		{5, 5},
		{39, 3},
		{871, 7},
		{5_642, 8},
		{93_201, 6},
		{123_456_789, 9},
		{2_147_483_647, 1}, // max int32
		{-1, -1},
		{-2_147_483_648, -2_147_483_648}, // min int32
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d", tc.n), func(t *testing.T) {
			if got := DigitalRoot(tc.n); got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}
