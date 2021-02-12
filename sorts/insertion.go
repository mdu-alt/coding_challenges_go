package sorts

// Insertion sorts a slice of integers using insertion sort algorithm.
func Insertion(slice []int) {
	for i := 1; i < len(slice); i++ {
		x := slice[i]
		j := i - 1

		// shift elements to the right...
		for j >= 0 && slice[j] > x {
			slice[j+1] = slice[j]
			j--
		}

		// ...then assign x in-place
		slice[j+1] = x
	}
}
