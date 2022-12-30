package numbers

// DigitalRoot computes the digital root of n.
//
// The digital root of a positive integer is the recursive sum of its digits,
// until a single digit remains. Returns n if n is negative.
func DigitalRoot(n int) int {
	var m int

	for n > 9 {
		for n != 0 {
			m += n % 10
			n /= 10
		}

		n, m = m, 0
	}

	return n
}
