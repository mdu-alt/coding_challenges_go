package numbers

import "strings"

// RomanToDecimal converts a string of roman numerals to a decimal number. It
// returns -1 in case the input contains:
//  - invalid symbols (e.g. 'A', '3', '$')
//  - too many symbols in a row (e.g. "IIII", "VV")
//  - invalid subtract cases (e.g. "IIX", "IXI", "VL")
func RomanToDecimal(roman string) int {
	const (
		I = 1 << iota // 1
		V             // 5
		X             // 10
		L             // 50
		C             // 100
		D             // 500
		M             // 1000
	)

	toBit := map[rune]int{
		'I': I,
		'V': V,
		'X': X,
		'L': L,
		'C': C,
		'D': D,
		'M': M,
	}

	toDecimal := map[int]int{
		I: 1,
		V: 5,
		X: 10,
		L: 50,
		C: 100,
		D: 500,
		M: 1000,
	}

	var (
		max            int = M << 1
		previous       int = M
		decimal, accum int
	)

	for _, r := range roman {
		current, ok := toBit[r]

		// Invalid cases: "A" / "3" / "$" / "VV" / "VIV" / "IXI" / "CMC"
		if !ok || current >= max {
			return -1
		}

		// Invalid cases: "IIX" / "CCM"
		if accum == 2 && current > previous {
			return -1
		}

		// Invalid cases: "IIII" / "CCCC"
		if current == previous {
			if accum++; accum > 3 {
				return -1
			}
		} else {
			accum = 1
		}

		if previous&(I|X|C) != 0 && current > previous && current <= (previous<<2) {
			// Cases: "IV" / "CM"
			decimal += toDecimal[current] - 2*toDecimal[previous]

			max = previous
		} else {
			// Case: "VI" / "CL"
			decimal += toDecimal[current]

			if current&(V|L|D) != 0 {
				max = current
			}
		}

		previous = current
	}

	return decimal
}

// DecimalToRoman converts a decimal number to a string of roman numerals.
//
//   Constraints:
//       - 3999 >= decimal >= 1
func DecimalToRoman(decimal int) string {
	symbols := []struct {
		value   int
		letters string
	}{
		{1000, "M"},
		{900, "CM"}, {500, "D"}, {400, "CD"}, {100, "C"},
		{90, "XC"}, {50, "L"}, {40, "XL"}, {10, "X"},
		{9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	}

	var b strings.Builder

	for _, pair := range symbols {
		quotient := decimal / pair.value
		b.WriteString(strings.Repeat(pair.letters, quotient))
		decimal -= pair.value * quotient
	}

	return b.String()
}
