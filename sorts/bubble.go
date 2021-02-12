package sorts

// Bubble sorts a slice of integers using bubble sort algorithm.
func Bubble(slice []int) {
	for len(slice) > 1 {
		up := 0

		for i := 1; i < len(slice); i++ {
			if slice[i-1] > slice[i] {
				slice[i-1], slice[i] = slice[i], slice[i-1]
				up = i
			}
		}

		// up == 0: no swaps happened, slice is sorted
		//    != 0: slice[up:] is sorted, don't loop beyond next time
		slice = slice[:up]
	}
}
