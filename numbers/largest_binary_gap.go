package numbers

// LargestBinaryGap finds the length of the largest binary gap of a positive
// integer. It returns -1 if n is negative.
//
// A binary gap of a positive integer n is any maximal sequence of consecutive
// zeros that is surrounded by ones in the binary representation of n.
func LargestBinaryGap(n int) int {
	if n < 0 {
		return -1
	}

	var (
		hasBit1   bool
		gap, accu int
	)

	for n != 0 {
		if n&1 == 0 && hasBit1 {
			accu++
		} else if n&1 == 1 {
			hasBit1 = true
			gap = max(gap, accu)
			accu = 0
		}

		n >>= 1
	}

	return gap
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
