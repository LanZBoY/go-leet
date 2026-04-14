package p0347_top_k_frequent_elements

func topKFrequent(nums []int, k int) []int {
	// 建立一個map
	countMap := make(map[int]int)

	for _, num := range nums{
		countMap[num]++
	}

	buckets := make([][]int, len(nums) + 1)

	for num, count :=  range countMap{
		buckets[count] = append(buckets[count], num)
	}

	result := make([]int, 0, k)

	for i := len(buckets) - 1; i>=0 && len(result) != k; i--{
		result = append(result, buckets[i]...)
	}

	return result
}
