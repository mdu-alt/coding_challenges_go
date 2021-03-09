package numbers

// LargestBinaryGap returns the length of the largest binary gap of n.
//
// A binary gap of a positive integer is any maximal sequence of consecutive
// zeros that is surrounded by ones in its binary representation.
//
//   Constraints:
//       - n >= 0
func LargestBinaryGap(n int) int {
	var (
		hasBit1    bool
		gap, accum int
	)

	for n != 0 {
		if n&1 == 0 && hasBit1 {
			accum++
		} else if n&1 == 1 {
			hasBit1 = true
			gap = max(gap, accum)
			accum = 0
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
