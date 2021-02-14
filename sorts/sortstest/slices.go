package sortstest

// Slices for testing purposes. Defined as functions to prevent test suites
// from modifying the slices.
var (
	S0 = func() []int { return []int{} }

	S1 = func() []int { return []int{7} }
	S2 = func() []int { return []int{3, 1} }
	S3 = func() []int { return []int{3, 5, 4} }

	S10 = func() []int {
		return []int{6, 5, 13, 11, 2, 12, 3, 4, 7, 10}
	}
	S20 = func() []int {
		return []int{30, 13, 22, 15, 23, 29, 28, 1, 20, 11, 27, 24, 26, 25, 10, 9, 17, 7, 6, 4}
	}
	S30 = func() []int {
		return []int{29, 43, 4, 10, 8, 40, 13, 47, 22, 25, 46, 14, 36, 41, 45, 23, 33, 15, 28, 34, 49, 30, 7, 17, 37, 18, 16, 21, 42, 6}
	}
)
