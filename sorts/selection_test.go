package sorts_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mdu-alt/coding_challenges_go/sorts"
	"github.com/mdu-alt/coding_challenges_go/sorts/sortstest"
)

func TestSelection(t *testing.T) {
	testCases := sortstest.Slices

	for _, tc := range testCases {
		var (
			got  = make([]int, len(tc))
			want = make([]int, len(tc))
		)

		copy(got, tc)
		copy(want, tc)

		sort.Ints(want)

		t.Run(fmt.Sprintf("%v", tc), func(t *testing.T) {
			if sorts.Selection(got); !cmp.Equal(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
