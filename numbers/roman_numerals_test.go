package numbers_test

import (
	"fmt"
	"testing"

	"github.com/mdu-alt/coding_challenges_go/numbers"
)

func ExampleRomanToDecimal() {
	fmt.Println(numbers.RomanToDecimal("XLII"))
	// Output:
	// 42
}

func TestRomanToDecimal(t *testing.T) {
	testCases := []struct {
		roman string
		want  int
	}{
		// single symbol
		{"I", 1},
		{"V", 5},
		{"D", 500},
		{"M", 1000},

		// many symbols
		{"VIII", 8},
		{"XXII", 22},
		{"DCLVI", 656},
		{"MCLXVII", 1167},

		// subtract
		{"IV", 4},
		{"XIX", 19},
		{"XLII", 42},
		{"MMCM", 2900},

		// invalid
		{"A", -1},
		{"IIII", -1},
		{"VV", -1},
		{"IIX", -1},
		{"IXI", -1},
		{"VX", -1},
		{"VIV", -1},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s", tc.roman), func(t *testing.T) {
			if got := numbers.RomanToDecimal(tc.roman); got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func ExampleDecimalToRoman() {
	fmt.Println(numbers.DecimalToRoman(42))
	// Output:
	// XLII
}

func TestDecimalToRoman(t *testing.T) {
	testCases := []struct {
		decimal int
		want    string
	}{
		{1, "I"},
		{10, "X"},
		{176, "CLXXVI"},
		{345, "CCCXLV"},
		{2489, "MMCDLXXXIX"},
		{3888, "MMMDCCCLXXXVIII"},
		{3999, "MMMCMXCIX"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d", tc.decimal), func(t *testing.T) {
			if got := numbers.DecimalToRoman(tc.decimal); got != tc.want {
				t.Errorf("got %s, want %s", got, tc.want)
			}
		})
	}
}
