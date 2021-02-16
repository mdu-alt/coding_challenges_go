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

// DecimalToRoman converts a decimal number to a string of roman numerals. It
// returns an empty string if decimal is out of the range [1..3999].
func DecimalToRoman(decimal int) string {
	if decimal < 1 || decimal > 3999 {
		return ""
	}

	symbols := strings.Split("IVXLCDM", "")

	var (
		offset int
		b      strings.Builder
	)

	// Each iteration, the loop works on each digit of the number (starting
	// from the unit) and a limited set of symbols is used each time: "IVX",
	// "XLC" and "CDM".
	for decimal != 0 {
		digit := decimal % 10

		switch {
		case 0 < digit && digit < 4:
			b.WriteString(strings.Repeat(symbols[offset], digit))
		case digit == 4:
			b.WriteString(symbols[offset+1] + symbols[offset])
		case 4 < digit && digit < 9:
			b.WriteString(strings.Repeat(symbols[offset], digit-5) + symbols[offset+1])
		case digit == 9:
			b.WriteString(symbols[offset+2] + symbols[offset])
		}

		decimal /= 10
		offset += 2
	}

	roman := []rune(b.String())

	for i, j := 0, len(roman)-1; i < j; i, j = i+1, j-1 {
		roman[i], roman[j] = roman[j], roman[i]
	}

	return string(roman)
}
