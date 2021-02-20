// Package sortstest defines testing-specific slices.
package sortstest

// Slices defines various testing slices.
var Slices = [][]int{
	// empty
	{},

	// all equal
	{3, 3, 3, 3, 3, 3, 3, 3, 3, 3},

	// increasing/decreasing
	{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},

	// any
	{5},
	{-1, 4, -1, -2, 3},
	{-4, 9, 3, -2, 5, 6, 0, 0, 7, -1},
	{6, 6, -8, -20, 4, -6, -10, -16, 11, -4, 10, -1, 5, 12, 16, 7, 15, 4, -11, 17},
}
