package arrays_test

import (
	"fmt"
	"testing"

	"github.com/mdu-alt/coding_challenges_go/arrays"
)

func ExampleLongestSubstring() {
	fmt.Println(arrays.LongestSubstring("Hello, World!"))
	// Output:
	// 8
}

func TestLongestSubstring(t *testing.T) {
	testCases := []struct {
		name string
		s    string
		want int
	}{
		{"empty", "", 0},

		{"distinct", "abcdefghij", 10},
		{"distinct", "0123456789", 10},
		{"distinct", "!@#$%^&*()", 10},

		{"repeating", "aaaaaaaaa", 1},
		{"repeating", "abcdeabcd", 5},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s:%q", tc.name, tc.s), func(t *testing.T) {
			if got := arrays.LongestSubstring(tc.s); got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}
