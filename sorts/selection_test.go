package sorts_test

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/mdu-alt/coding_challenges_go/sorts"
)

func TestSelection(t *testing.T) {
	testMatrix := map[string][]struct {
		slice []int
		want  []int
	}{
		"empty": {
			{slice0, make([]int, len(slice0))},
		},
		"any": {
			{slice10, make([]int, len(slice10))},
			{slice50, make([]int, len(slice50))},
			{slice100, make([]int, len(slice100))},
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
