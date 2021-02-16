package binary_test

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	. "github.com/mdu-alt/coding_challenges_go/trees/binary"
)

var (
	bstLeft = &Tree{
		Value: 5,
		Left: &Tree{
			Value: 4,
			Left: &Tree{
				Value: 3,
				Left: &Tree{
					Value: 2,
					Left: &Tree{
						Value: 1,
					},
				},
			},
		},
	}
	bstRight = &Tree{
		Value: 1,
		Right: &Tree{
			Value: 2,
			Right: &Tree{
				Value: 3,
				Right: &Tree{
					Value: 4,
					Right: &Tree{
						Value: 5,
					},
				},
			},
		},
	}
)

func TestNewBST(t *testing.T) {
	testCases := []struct {
		slice []int
		want  *Tree
	}{
		// empty
		{[]int{}, nil},

		// left/right
		{[]int{5, 4, 3, 2, 1}, bstLeft},
		{[]int{1, 2, 3, 4, 5}, bstRight},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.slice), func(t *testing.T) {
			if got := NewBST(tc.slice); !cmp.Equal(got, tc.want) {
				t.Errorf("diff -got +want\n%v", cmp.Diff(got, tc.want))
			}
		})
	}
}
