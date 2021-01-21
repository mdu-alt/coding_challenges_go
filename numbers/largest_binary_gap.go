package numbers

import "math"

func largestBinaryGap(n int) int {
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
			gap = int(math.Max(float64(accu), float64(gap)))
			accu = 0
		}

		n >>= 1
	}

	return gap
}
