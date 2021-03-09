package numbers

// DigitalRoot computes the digital root of n.
//
// The digital root of a positive integer is the recursive sum of its digits,
// until a single digit remains.
//
//   Constraints:
//       - n >= 0
func DigitalRoot(n int) int {
	for n > 9 {
		var m int

		for n != 0 {
			m += n % 10
			n /= 10
		}

		n = m
	}

	return n
}
