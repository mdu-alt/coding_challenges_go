package arrays

// BinarySearch returns the index of the first instance of n in a sorted slice.
// It returns -1 if n is not present in s.
func BinarySearch(slice []int, n int) int {
	if len(slice) == 0 {
		return -1
	}

	i, j := 0, len(slice)-1

	for i < j {
		m := int(uint(i+j) >> 1) // avoid overflow

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
