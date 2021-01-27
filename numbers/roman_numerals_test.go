package numbers_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/mdu-alt/coding_challenges_go/numbers"
)

func ExampleRomanToDecimal() {
	fmt.Println(numbers.RomanToDecimal("XLII"))
	// Output: 42
}

func TestRomanToDecimal(t *testing.T) {
	testMatrix := map[string][]struct {
		roman string
		want  int
	}{
		"single symbol": {
			{"I", 1},
			{"V", 5},
			{"D", 500},
			{"M", 1000},
		},
		"many symbols": {
			{"VIII", 8},
			{"XXII", 22},
			{"DCLVI", 656},
			{"MCLXVII", 1167},
		},
		"subtract": {
			{"IV", 4},
			{"XIX", 19},
			{"XLII", 42},
			{"MMCM", 2900},
		},
		"invalid": {
			{"A", -1},
			{"IIII", -1},
			{"VV", -1},
			{"IIX", -1},
			{"IXI", -1},
			{"VX", -1},
			{"VIV", -1},
		},
	}

	for title, testCases := range testMatrix {
		for _, test := range testCases {
			name := fmt.Sprintf("%s{%q|%d}", title, test.roman, test.want)

			t.Run(name, func(t *testing.T) {
				if got := numbers.RomanToDecimal(test.roman); got != test.want {
					t.Errorf("expected: %d, got: %d", test.want, got)
				}
			})
		}
	}
}

func ExampleDecimalToRoman() {
	fmt.Println(numbers.DecimalToRoman(42))
	// Output: XLII
}

func TestDecimalToRoman(t *testing.T) {
	testMatrix := map[string][]struct {
		decimal int
		want    string
	}{
		"valid": {
			{1, "I"},
			{176, "CLXXVI"},
			{345, "CCCXLV"},
			{2489, "MMCDLXXXIX"},
			{3888, "MMMDCCCLXXXVIII"},
			{3999, "MMMCMXCIX"},
		},
		"invalid": {
			{math.MinInt32, ""},
			{0, ""},
			{4000, ""},
			{math.MaxInt32, ""},
		},
	}

	for title, testCases := range testMatrix {
		for _, test := range testCases {
			name := fmt.Sprintf("%s{%d|%q}", title, test.decimal, test.want)

			t.Run(name, func(t *testing.T) {
				if got := numbers.DecimalToRoman(test.decimal); got != test.want {
					t.Errorf("expected: %s, got: %s", test.want, got)
				}
			})
		}
	}
}
