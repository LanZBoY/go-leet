package p0001_two_sum

func twoSum(nums []int, target int) []int {

	// Create a map which is Remain to Index
	rMap := make(map[int]int)

	for i, num := range nums {

		if findIndex, ok := rMap[num]; ok {
			return []int{findIndex, i}
		}

		rMap[target-num] = i
	}

	return nil
}
