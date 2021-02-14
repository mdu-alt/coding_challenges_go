package sorts

// Quick sorts a slice of integers using quick sort algorithm.
func Quick(slice []int) {
	if len(slice) < 2 {
		return
	}

	i, j := 0, len(slice)-1
	pivot := slice[len(slice)/2]

	for i < j {
		for slice[i] < pivot {
			i++
		}
		for slice[j] > pivot {
			j--
		}

		slice[i], slice[j] = slice[j], slice[i]
	}

	Quick(slice[:i])
	Quick(slice[i+1:])
}
