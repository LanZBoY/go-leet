package p0128_longest_consecutive_sequence

func longestConsecutive(nums []int) int {
	numSet := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		numSet[num] = struct{}{}
	}

	maxLen := 0
	for num := range numSet {
		if _, ok := numSet[num-1]; ok {
			continue // num 不是連續序列起點
		}
		curLen := 1
		for {
			if _, ok := numSet[num+curLen]; !ok {
				break
			}
			curLen++
		}
		if curLen > maxLen {
			maxLen = curLen
		}
	}

	return maxLen
}
