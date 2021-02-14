package sorts_test

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/mdu-alt/coding_challenges_go/sorts"
	"github.com/mdu-alt/coding_challenges_go/sorts/sortstest"
)

func TestSelection(t *testing.T) {
	testMatrix := map[string][]struct {
		slice []int
		want  []int
	}{
		"empty": {
			{sortstest.S0(), make([]int, len(sortstest.S0()))},
		},
		"any": {
			{sortstest.S1(), make([]int, len(sortstest.S1()))},
			{sortstest.S2(), make([]int, len(sortstest.S2()))},
			{sortstest.S3(), make([]int, len(sortstest.S3()))},
			{sortstest.S10(), make([]int, len(sortstest.S10()))},
			{sortstest.S20(), make([]int, len(sortstest.S20()))},
			{sortstest.S30(), make([]int, len(sortstest.S30()))},
		},
	}

	for title, testCases := range testMatrix {
		for _, test := range testCases {
			name := fmt.Sprintf("%s{%v}", title, test.slice)

			t.Run(name, func(t *testing.T) {
				copy(test.want, test.slice)
				sort.Ints(test.want)

				if sorts.Selection(test.slice); !reflect.DeepEqual(test.slice, test.want) {
					t.Errorf("expected: %v, got: %v", test.want, test.slice)
				}
			})
		}
	}
}
