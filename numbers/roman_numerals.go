package numbers

import (
	"strings"
)

const (
	symbolI = 1 << iota // 1
	symbolV             // 5
	symbolX             // 10
	symbolL             // 50
	symbolC             // 100
	symbolD             // 500
	symbolM             // 1000
)

var map2bit = map[byte]int{
	'M': symbolM,
	'D': symbolD,
	'C': symbolC,
	'L': symbolL,
	'X': symbolX,
	'V': symbolV,
	'I': symbolI,
}

var map2decimal = map[int]int{
	symbolM: 1000,
	symbolD: 500,
	symbolC: 100,
	symbolL: 50,
	symbolX: 10,
	symbolV: 5,
	symbolI: 1,
}

// var toto = map[int][]int]{
// 	symbolM: 1000,
// 	symbolD: 500,
// 	symbolC: 100,
// 	symbolL: 50,
// 	symbolX: 10,
// 	symbolV: 5,
// 	symbolI: 1,
// }

func romanToDecimal(roman string) int {
	var (
		accu      []int
		decimal   int
		prev, min int
	)

	for i := len(roman) - 1; i >= 0; i-- {
		curr := map2bit[roman[i]]

		// Invalid cases: "IIX", "IXI"
		if curr == 0 || curr < min {
			return -1
		}

		if len(accu) > 0 && accu[len(accu)-1] == curr {
			accu = append(accu, curr)

			// Invalid case: "IIII"
			if len(accu) > 3 {
				return -1
			}
		} else {
			accu = []int{curr}
		}

		// Invalid cases: "VV", "LL", "DD"
		if curr&prev&(symbolV|symbolL|symbolD) != 0 {
			return -1
		}

		if curr&(symbolI|symbolX|symbolC) != 0 && curr < prev && prev <= (curr<<2) {
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
			// Invalid cases: "VX", "IL"
			return -1
		}

		prev = curr
	}

	return decimal
}

func decimalToRoman(decimal int) string {
	if decimal < 1 || decimal > 3999 {
		return ""
	}

	symbols := []string{"I", "V", "X", "L", "C", "D", "M"}

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
