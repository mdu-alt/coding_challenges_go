package sorts

// Quick sorts a slice of integers using quick sort algorithm.
func Quick(slice []int) {
	if len(slice) < 2 {
		return
	}

	// Hoare partition scheme.
	i, j := -1, len(slice)
	pivot := slice[(len(slice)-1)/2]

	for {
		i++
		for slice[i] < pivot {
			i++
		}

		j--
		for slice[j] > pivot {
			j--
		}

		if i >= j {
			break
		}

		slice[i], slice[j] = slice[j], slice[i]
	}

	Quick(slice[:j+1])
	Quick(slice[j+1:])
}
