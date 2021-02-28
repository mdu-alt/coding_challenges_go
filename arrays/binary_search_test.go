package arrays_test

import (
	"fmt"
	"testing"

	"github.com/mdu-alt/coding_challenges_go/arrays"
)

func ExampleBinarySearch() {
	fmt.Println(arrays.BinarySearch([]int{2, 5, 8, 11, 17}, 11))
	// Output:
	// 3
}

func TestBinarySearch(t *testing.T) {
	slice := []int{0, 2, 3, 4, 5, 5, 8, 10, 13, 14}
	testCases := []struct {
		name  string
		slice []int
		n     int
		want  int
	}{
		{"empty", []int{}, 0, -1},

		{"found", slice, 0, 0},
		{"found", slice, 3, 2},
		{"found", slice, 5, 4},
		{"found", slice, 8, 6},
		{"found", slice, 14, 9},

		{"not found", slice, 6, -1},
		{"not found", slice, -1, -1},
		{"not found", slice, 15, -1},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s:%d", tc.name, tc.n), func(t *testing.T) {
			if got := arrays.BinarySearch(tc.slice, tc.n); got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}
