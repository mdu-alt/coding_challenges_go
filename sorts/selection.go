package sorts

// Selection sorts a slice of integers using selection sort algorithm.
func Selection(slice []int) {
	for i := range slice {
		m := i

		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[m] {
				m = j
			}
		}

		slice[i], slice[m] = slice[m], slice[i]
	}
}
