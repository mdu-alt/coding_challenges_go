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
			{"X", 10},
			{"D", 500},
		},
		"many symbols": {
			{"VI", 6},
			{"XXII", 22},
			{"MCLXVII", 1167},
		},
		"subtract": {
			{"IV", 4},
			{"XIX", 19},
			{"MMCM", 2900},
		},
		"invalid": {
			{"A", -1},
			{"IIII", -1},
			{"VV", -1},
			{"IIX", -1},
			{"IXI", -1},
			{"VX", -1},
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
