package singly_test

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mdu-alt/coding_challenges_go/lists/singly"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		list []int
		want []int
	}{
		{nil, nil},

		{[]int{1}, []int{1}},
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{5, 4, 3, 2, 1, 0, -1, -2, -3, -4}, []int{-4, -3, -2, -1, 0, 1, 2, 3, 4, 5}},
	}

	for _, tc := range testCases {
		list := singly.New(tc.list)
		want := singly.New(tc.want)

		t.Run(fmt.Sprintf("%v", tc.list), func(t *testing.T) {
			if got := singly.Reverse(list); !cmp.Equal(got, want) {
				t.Errorf("diff -got +want\n%v", cmp.Diff(got, want))
			}
		})
	}
}
