package numbers

func digitalRoot(n int) int {
	if n < 0 {
		return -1
	}

	for n/10 != 0 {
		l := 0

		for n != 0 {
			l += n % 10
			n /= 10
		}

		n = l
	}

	return n
}
