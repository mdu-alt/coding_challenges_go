package strings_test

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mdu-alt/coding_challenges_go/strings"
)

func ExampleAllPalindromes() {
	fmt.Println(strings.AllPalindromes("Racecar wow"))
	// Output:
	// [cec aceca wow]
}

func TestAllPalindromes(t *testing.T) {
	testCases := []struct {
		name string
		s    string
		want []string
	}{
		{"odd", "a", nil},
		{"odd", "abc", nil},
		{"odd", "abc def ghi", nil},
		{"odd", "aba", []string{"aba"}},
		{"odd", "abcba", []string{"bcb", "abcba"}},

		{"even", "", nil},
		{"even", "ab", nil},
		{"even", "abcd efgh ", nil},
		{"even", "bb", []string{"bb"}},
		{"even", "abba", []string{"bb", "abba"}},

		{"mix", "aba bb cd", []string{"aba", "bb", " bb "}},
		{"mix", "racecar 12321 anna", []string{"cec", "aceca", "racecar", "232", "12321", " 12321 ", "nn", "anna"}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%q:%q", tc.name, tc.s), func(t *testing.T) {
			if got := strings.AllPalindromes(tc.s); !cmp.Equal(got, tc.want) {
				t.Errorf("diff -got +want\n%v", cmp.Diff(got, tc.want))
			}
		})
	}
}
