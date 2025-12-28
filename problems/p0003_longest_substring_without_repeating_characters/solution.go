package p0003_longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	nonRepeatSet := make(map[byte]struct{})
	maxSize := 0
	i := 0

	for j := 0; j < len(s); j++ {

		for {
			if _, ok := nonRepeatSet[s[j]]; !ok {
				break
			}
			delete(nonRepeatSet, s[i])
			i++
		}

		nonRepeatSet[s[j]] = struct{}{}

		maxSize = max(maxSize, len(nonRepeatSet))
	}
	return maxSize
}
