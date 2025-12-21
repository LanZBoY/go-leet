package p0347_top_k_frequent_elements

import "slices"

func topKFrequent(nums []int, k int) []int {
	// 建立一個map
	cMap := make(map[int]int)

	for _, num := range nums {
		cMap[num]++
	}

	type NumCount struct {
		num   int
		count int
	}

	ncs := make([]NumCount, 0)

	for k, v := range cMap {
		ncs = append(ncs, NumCount{
			num:   k,
			count: v,
		})
	}

	slices.SortFunc(ncs, func(a NumCount, b NumCount) int {
		return b.count - a.count
	})

	result := make([]int, k)

	for i := range k {
		result[i] = ncs[i].num
	}

	return result
}
