package numbers

// DigitalRoot repeatedly sums all digits of a positive integer n until a
// single digit remains. It returns -1 if n is negative.
func DigitalRoot(n int) int {
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
