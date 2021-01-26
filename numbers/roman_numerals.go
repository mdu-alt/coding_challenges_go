package numbers

import "strings"

// RomanToDecimal converts a string of roman numerals to a decimal number.
//
// It returns -1 in case of an invalid input, such as:
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

	map2bit := map[byte]int{
		'I': I,
		'V': V,
		'X': X,
		'L': L,
		'C': C,
		'D': D,
		'M': M,
	}

	map2decimal := map[int]int{
		I: 1,
		V: 5,
		X: 10,
		L: 50,
		C: 100,
		D: 500,
		M: 1000,
	}

	var (
		accu      []int
		decimal   int
		prev, min int
	)

	for i := len(roman) - 1; i >= 0; i-- {
		curr, ok := map2bit[roman[i]]

		// Invalid cases: "A3$" / "IIX" / "IXI"
		if !ok || curr < min {
			return -1
		}

		// Invalid case: "IIII"
		if len(accu) > 0 && accu[len(accu)-1] == curr {
			if accu = append(accu, curr); len(accu) > 3 {
				return -1
			}
		} else {
			accu = []int{curr}
		}

		// Invalid cases: "VV" / "LL" / "DD"
		if curr&prev&(V|L|D) != 0 {
			return -1
		}

		if curr&(I|X|C) != 0 && curr < prev && prev <= (curr<<2) {
			// Case: "IV"
			decimal -= map2decimal[curr]
			min = prev
		} else if curr >= prev {
			// Case: "VI"
			decimal += map2decimal[curr]

			if prev != 0 {
				min = curr
			}
		} else {
			// Invalid cases: "VX" / "IL"
			return -1
		}

		prev = curr
	}

	return decimal
}

// DecimalToRoman converts a decimal number to a string of roman numerals.
// It returns an empty string if decimal is out of the range [1..3999].
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

	roman := []byte(b.String())

	for i, j := 0, len(roman)-1; i < j; i, j = i+1, j-1 {
		roman[i], roman[j] = roman[j], roman[i]
	}

	return string(roman)
}
