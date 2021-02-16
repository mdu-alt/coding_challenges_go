package sequences

import "sort"

// BinarySearch returns the index of the first instance of n in a sorted slice.
// It returns -1 if n is not present in s or if s is not sorted in increasing
// order.
func BinarySearch(slice []int, n int) int {
	if len(slice) == 0 || !sort.IntsAreSorted(slice) {
		return -1
	}

	i, j := 0, len(slice)-1

	for i < j {
		m := i + (j-i)/2 // avoid overflow

		if n > slice[m] {
			i = m + 1
		} else {
			j = m
		}
	}

	if slice[i] == n {
		return i
	}

	return -1
}
