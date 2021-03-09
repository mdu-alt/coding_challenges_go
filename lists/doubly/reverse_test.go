package doubly_test

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mdu-alt/coding_challenges_go/lists/doubly"
)

func TestReverse(t *testing.T) {
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
		list := doubly.New(tc.list)
		want := doubly.New(tc.want)

		t.Run(fmt.Sprintf("%s:%v", tc.name, tc.list), func(t *testing.T) {
			if got := doubly.Reverse(list); !cmp.Equal(got, want) {
				t.Errorf("diff -got +want\n%v", cmp.Diff(got, want))
			}
		})
	}
}
