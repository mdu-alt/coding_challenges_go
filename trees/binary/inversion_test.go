package binary_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	. "github.com/mdu-alt/coding_challenges_go/trees/binary"
	"github.com/mdu-alt/coding_challenges_go/trees/binary/binarytest"
)

func TestInvert(t *testing.T) {
	testCases := []struct {
		name string
		tree *Tree
		want *Tree
	}{
		{"empty", binarytest.Empty, binarytest.Empty},
		{"simple", binarytest.Simple, binarytest.SimpleInvert},
		{"left", binarytest.Left, binarytest.LeftInvert},
		{"right", binarytest.Right, binarytest.RightInvert},
		{"any", binarytest.Any, binarytest.AnyInverted},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.tree
			if Invert(got); !cmp.Equal(got, tc.want) {
				t.Errorf("diff -got +want\n%v", cmp.Diff(got, tc.want))
			}
		})
	}
}
