package sequences_test

import (
	"fmt"
	"testing"

	"github.com/mdu-alt/coding_challenges_go/sequences"
)

func ExampleBinarySearch() {
	fmt.Println(sequences.BinarySearch([]int{2, 5, 8, 11, 17}, 11))
	// Output:
	// 3
}

func TestBinarySearch(t *testing.T) {
	slice := []int{0, 2, 3, 4, 5, 5, 8, 10, 13, 14}
	testCases := []struct {
		slice []int
		n     int
		want  int
	}{
		// empty
		{[]int{}, 0, -1},

		// found
		{slice, 0, 0},
		{slice, 3, 2},
		{slice, 5, 4},
		{slice, 8, 6},
		{slice, 14, 9},

		// not found
		{slice, 6, -1},
		{slice, -1, -1},
		{slice, 15, -1},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%v,%d", tc.slice, tc.n)
		t.Run(name, func(t *testing.T) {
			got := sequences.BinarySearch(tc.slice, tc.n)
			if got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}
