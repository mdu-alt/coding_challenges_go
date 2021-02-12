package sorts

// Selection sorts a slice of integers using selection sort algorithm.
func Selection(slice []int) {
	for i := range slice {
		j := i

		for k := range slice[i+1:] {
			if slice[k+i+1] < slice[j] {
				j = k + i + 1
			}
		}

		if i != j {
			slice[i], slice[j] = slice[j], slice[i]
		}
	}
}
