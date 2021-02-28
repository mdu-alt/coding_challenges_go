package arrays

// LongestSubstring finds the length of the longest substring without repeating
// characters.
func LongestSubstring(s string) int {
	var (
		hashmap = make(map[rune]int)
		i, m    int
	)

	for j, r := range s {
		if _, ok := hashmap[r]; ok {
			i = max(i, hashmap[r])
		}

		m = max(m, j-i+1)

		hashmap[r] = j + 1
	}

	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
