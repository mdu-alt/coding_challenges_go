package numbers_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/mdu-alt/coding_challenges_go/numbers"
)

func ExampleLargestBinaryGap() {
	// 25 = 0b11001 -> gap = 2
	fmt.Println(numbers.LargestBinaryGap(25))
	// Output:
	// 2
}

func TestLargestBinaryGap(t *testing.T) {
	testCases := []struct {
		n    int
		want int
	}{
		// any gaps
		{5, 1},          //          5 = 0b101
		{9, 2},          //          9 = 0b1001
		{2546, 2},       //       2546 = 0b100111110010
		{1376796946, 5}, // 1376796946 = 0b1010010000100000100000100010010

		// no gaps
		{1, 0},  //  1 = 0b1
		{3, 0},  //  3 = 0b11
		{24, 0}, // 24 = 0b11000
		{32, 0}, // 32 = 0b100000

		// min/max bounds
		{0, 0},             //   0 = 0b0
		{math.MaxInt32, 0}, // Max = 0b1111111111111111111111111111111

		// negative
		{-1, -1},
		{math.MinInt32, -1},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d", tc.n), func(t *testing.T) {
			if got := numbers.LargestBinaryGap(tc.n); got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}
