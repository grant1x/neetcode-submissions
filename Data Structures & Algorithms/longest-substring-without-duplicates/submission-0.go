func lengthOfLongestSubstring(s string) int {
	maxT := 0
	seen := make(map[rune]int)

	l := 0
	for r := range s {
		char := rune(s[r])

		if oldIndex, exists := seen[char]; exists && oldIndex >= l {
			l = oldIndex + 1
		}

      seen[char] = r

		maxT = max(maxT, r-l+1)
	}
	return maxT
}