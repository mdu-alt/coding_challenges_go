// Package sortstest defines testing slices.
package sortstest

// Various testing slices.
var (
	Slice0 []int

	Slice1  = []int{5}
	Slice5  = []int{-1, 4, -1, -2, 3}
	Slice10 = []int{-4, 9, 3, -2, 5, 6, 0, 0, 7, -1}

	SliceEqual10 = []int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3}
	SliceEqual20 = []int{-7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7}

	SliceIncreasing10 = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	SliceDecreasing10 = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
)
