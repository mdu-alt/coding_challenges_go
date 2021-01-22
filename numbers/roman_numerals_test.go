package numbers

import (
	"fmt"
	"math"
	"testing"
)

func TestRomanToDecimal(t *testing.T) {
	testMatrix := map[string][]struct {
		roman    string
		expected int
	}{
		"singleSymbols": {
			{"I", 1},
			{"X", 10},
			{"D", 500},
		},
		"manySymbols": {
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

	for name, tests := range testMatrix {
		for _, test := range tests {
			testname := fmt.Sprintf("%s--{%s,%d}", name, test.roman, test.expected)

			t.Run(testname, func(t *testing.T) {
				if got := romanToDecimal(test.roman); got != test.expected {
					t.Errorf("expected: %d, got: %d", test.expected, got)
				}
			})
		}
	}
}

func TestDecimalToRoman(t *testing.T) {
	testMatrix := map[string][]struct {
		decimal  int
		expected string
	}{
		"valid": {
			{1, "I"},
			{45, "XLV"},
			{679, "DCLXXIX"},
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

	for name, tests := range testMatrix {
		for _, test := range tests {
			testname := fmt.Sprintf("%s--{%d,%s}", name, test.decimal, test.expected)

			t.Run(testname, func(t *testing.T) {
				if got := decimalToRoman(test.decimal); got != test.expected {
					t.Errorf("expected: %s, got: %s", test.expected, got)
				}
			})
		}
	}
}
