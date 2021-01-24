package sequences

import "sort"

// BinarySearch finds the smallest index of an integer e within a slice of
// sorted integers s. It returns -1 if e is not found or if s is empty.
func BinarySearch(s []int, e int) int {
	if len(s) == 0 {
		return -1
	}

	if !sort.IntsAreSorted(s) {
		sort.Ints(s) // safety sort
	}

	l, r := 0, len(s)-1

	for l < r {
		m := l + (r-l)/2 // avoid overflow

		if e <= s[m] {
			r = m
		} else {
			l = m + 1
		}
	}

	if s[l] == e {
		return l
	}

	return -1
}
