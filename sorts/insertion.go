package sorts

// Insertion sorts a slice of integers using insertion sort algorithm.
func Insertion(slice []int) {
	for i := 1; i < len(slice); i++ {
		x := slice[i]
		j := i - 1

		for j >= 0 && slice[j] > x {
			j--
		}

		// Shift elements to the right, then assign x in-place.
		copy(slice[j+2:i+1], slice[j+1:i])
		slice[j+1] = x
	}
}
