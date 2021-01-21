package numbers

import (
	"fmt"
	"math"
	"testing"
)

func TestRomanToDecimal(t *testing.T) {
	type testCase struct {
		roman    string
		expected int
	}

	testMatrix := map[string][]testCase{
		"valid": {
			{"I", 1},
			{"X", 10},
			{"D", 500},
		},
		"many_symbols": {
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
			testname := fmt.Sprintf("%v--{%v,%v}", name, test.roman, test.expected)

			t.Run(testname, func(t *testing.T) {
				if got := romanToDecimal(test.roman); got != test.expected {
					t.Errorf("expected: %v, got: %v", test.expected, got)
				}
			})
		}
	}
}

func TestDecimalToRoman(t *testing.T) {
	type testCase struct {
		decimal  int
		expected string
	}

	testMatrix := map[string][]testCase{
		"valid": {
			{1, "I"},
			{45, "XLV"},
			{679, "DCLXXIX"},
			{3888, "MMMDCCCLXXXVIII"},
			{3999, "MMMCMXCIX"},
		},
		"invalid": {
			{math.MinInt32, ""},
			{-1, ""},
			{0, ""},
			{4000, ""},
			{math.MaxInt32, ""},
		},
	}

	for name, tests := range testMatrix {
		for _, test := range tests {
			testname := fmt.Sprintf("%v--{%v,%v}", name, test.decimal, test.expected)

			t.Run(testname, func(t *testing.T) {
				if got := decimalToRoman(test.decimal); got != test.expected {
					t.Errorf("expected: %v, got: %v", test.expected, got)
				}
			})
		}
	}
}
