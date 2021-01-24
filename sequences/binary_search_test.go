package sequences

import "testing"

func TestBinarySearch(t *testing.T) {
	s := []int{0, 2, 3, 4, 5, 5, 8, 10, 13, 14}

	testMatrix := map[string][]struct {
		s    []int
		e    int
		want int
	}{
		"found": {
			{s, 0, 0},
			{s, 3, 2},
			{s, 5, 4},
			{s, 8, 6},
			{s, 14, 9},
		},
		"not found": {
			{s, 6, -1},
			{s, -1, -1},
			{s, 15, -1},
		},
		"empty slice": {
			{[]int{}, 0, -1},
		},
	}

	for name, testCases := range testMatrix {
		for _, test := range testCases {
			t.Run(name, func(t *testing.T) {
				if got := BinarySearch(test.s, test.e); got != test.want {
					t.Errorf("expected: %d, got: %d", test.want, got)
				}
			})
		}
	}
}
