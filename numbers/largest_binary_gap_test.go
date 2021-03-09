package numbers_test

import (
	"fmt"
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
		{5, 1},              //             5 = 0b101
		{9, 2},              //             9 = 0b1001
		{2_546, 2},          //         2_546 = 0b1001_1111_0010
		{220_267_537, 5},    //   220_267_537 = 0b1101_0010_0001_0000_0100_0001_0001
		{1_073_741_825, 29}, // 1_073_741_824 = 0b100_0000_0000_0000_0000_0000_0000_0001

		// no gaps
		{1, 0},      //      1 = 0b1
		{3, 0},      //      3 = 0b11
		{15, 0},     //     15 = 0b1111
		{65_280, 0}, // 65_280 = 0b1111_1111_0000_0000

		// min/max bounds
		{0, 0},             //             0 = 0b0
		{2_147_483_647, 0}, // 2_147_483_647 = 0b111_1111_1111_1111_1111_1111_1111_1111
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d", tc.n), func(t *testing.T) {
			if got := numbers.LargestBinaryGap(tc.n); got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}
