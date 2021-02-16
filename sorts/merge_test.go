package sorts_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mdu-alt/coding_challenges_go/sorts"
	"github.com/mdu-alt/coding_challenges_go/sorts/sortstest"
)

func TestMerge(t *testing.T) {
	testCases := [][]int{
		// empty
		sortstest.Slice0,

		// equal
		sortstest.SliceEqual10,
		sortstest.SliceEqual20,

		// increasing / decreasing
		sortstest.SliceIncreasing10,
		sortstest.SliceDecreasing10,

		// any
		sortstest.Slice1,
		sortstest.Slice5,
		sortstest.Slice10,
	}

	for _, tc := range testCases {
		var (
			got  = make([]int, len(tc))
			want = make([]int, len(tc))
		)

		copy(got, tc)
		copy(want, tc)

		sort.Ints(want)

		name := fmt.Sprintf("%v", tc)
		t.Run(name, func(t *testing.T) {
			sorts.Merge(got)
			if !cmp.Equal(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
