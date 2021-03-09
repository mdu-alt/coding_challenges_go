package arrays_test

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mdu-alt/coding_challenges_go/arrays"
)

func ExampleTwoSum() {
	fmt.Println(arrays.TwoSum([]int{1, 2, 3}, 5))
	// Output:
	// [1 2]
}

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		name string
		s    []int
		n    int
		want []int
	}{
		{"one pair", []int{4, 4}, 8, []int{0, 1}},
		{"one pair", []int{2, 7, 11, 15, 3}, 9, []int{0, 1}},
		{"one pair", []int{10, -9, 1, 2, 4}, 6, []int{3, 4}},

		{"no pair", []int{}, 5, nil},
		{"no pair", []int{5}, 5, nil},
		{"no pair", []int{1, 2, 5, -3, -6}, 5, nil},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s:%v:%d", tc.name, tc.s, tc.n), func(t *testing.T) {
			if got := arrays.TwoSum(tc.s, tc.n); !cmp.Equal(got, tc.want) {
				t.Errorf("diff -got +want\n%v", cmp.Diff(got, tc.want))
			}
		})
	}
}
