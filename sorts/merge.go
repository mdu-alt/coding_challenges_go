package sorts

// Merge sorts a slice of integers using merge sort algorithm.
func Merge(slice []int) {
	if len(slice) < 2 {
		return
	}

	pl := slice[:len(slice)/2]
	pr := slice[len(slice)/2:]

	Merge(pl)
	Merge(pr)

	sorted := make([]int, 0, len(pl)+len(pr))
	i, j := 0, 0

	for i < len(pl) && j < len(pr) {
		if pl[i] < pr[j] {
			sorted = append(sorted, pl[i])
			i++
		} else {
			sorted = append(sorted, pr[j])
			j++
		}
	}

	if i != len(pl) {
		sorted = append(sorted, pl[i:]...)
	} else if j != len(pr) {
		sorted = append(sorted, pr[j:]...)
	}

	copy(slice, sorted)
}
