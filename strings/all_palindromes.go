package strings

// AllPalindromes finds all palindromes (of length > 1) within ASCII string s.
func AllPalindromes(s string) []string {
	var allPalindromes []string

	for i := range s {
		loop(i-1, i+1, s, &allPalindromes)
		loop(i, i+1, s, &allPalindromes)
	}

	return allPalindromes
}

func loop(i, j int, s string, slice *[]string) {
	for ; i >= 0 && j < len(s); i, j = i-1, j+1 {
		if s[i] == s[j] {
			*slice = append(*slice, s[i:j+1])
		} else {
			break
		}
	}
}
