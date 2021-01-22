package algorithms

import (
	"math"
	"sort"
)

func binarySearch(s []int, e int) int {
	sort.Ints(s)

	l, r := 0, len(s)-1

	for l != r {
		m := int(math.Ceil(float64(l+r) / 2))

		if e < s[m] {
			r = m - 1
		} else {
			l = m
		}
	}

	if s[l] == e {
		return l
	}

	return -1
}
