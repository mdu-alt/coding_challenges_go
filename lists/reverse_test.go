package lists_test

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	. "github.com/mdu-alt/coding_challenges_go/lists"
	"github.com/mdu-alt/coding_challenges_go/lists/liststest"
)

func TestReverseSingly(t *testing.T) {
	testCases := []struct {
		name string
		list []int
		want []int
	}{
		{"empty", nil, nil},

		{"any", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"any", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"any", []int{3, 9, -2, 5, 0}, []int{0, 5, -2, 9, 3}},
	}

	for _, tc := range testCases {
		list := liststest.NewSingly(tc.list)
		want := liststest.NewSingly(tc.want)

		t.Run(fmt.Sprintf("%s:%v", tc.name, tc.list), func(t *testing.T) {
			if got := ReverseSingly(list); !cmp.Equal(got, want) {
				t.Errorf("diff -got +want\n%v", cmp.Diff(got, want))
			}
		})
	}
}

func TestReverseDoubly(t *testing.T) {
	testCases := []struct {
		name string
		list []int
		want []int
	}{
		{"empty", nil, nil},

		{"any", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"any", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"any", []int{3, 9, -2, 5, 0}, []int{0, 5, -2, 9, 3}},
	}

	for _, tc := range testCases {
		list := liststest.NewDoubly(tc.list)
		want := liststest.NewDoubly(tc.want)

		t.Run(fmt.Sprintf("%s:%v", tc.name, tc.list), func(t *testing.T) {
			if got := ReverseDoubly(list); !cmp.Equal(got, want) {
				t.Errorf("diff -got +want\n%v", cmp.Diff(got, want))
			}
		})
	}
}
