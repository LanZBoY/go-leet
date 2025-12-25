package p0287_find_the_duplicate_number

func findDuplicate(nums []int) int {
	return sln_4(nums)
}

// Floyd's Tortoise and Hare (Cycle Detection)
// Time: O(N), Space: O(1)
func sln_4(nums []int) int {
	// 1. Move tortoise and hare to find the meeting point
	tortoise := nums[0]
	hare := nums[0]

	for {
		tortoise = nums[tortoise]
		hare = nums[nums[hare]]
		if tortoise == hare {
			break
		}
	}

	// 2. Find the entrance to the cycle
	ptr1 := nums[0]
	ptr2 := tortoise
	for ptr1 != ptr2 {
		ptr1 = nums[ptr1]
		ptr2 = nums[ptr2]
	}

	return ptr1
}

// Binary Search
// Time: O(N log N), Space: O(1)
func sln_3(nums []int) int {
	left, right := 1, len(nums)-1

	for left < right {
		mid := (left + right) / 2
		count := 0

		// Count how many numbers are less than or equal to mid
		for _, num := range nums {
			if num <= mid {
				count++
			}
		}

		if count > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

// Hash Map
// Time: O(N), Space: O(N)
func sln_2(nums []int) int {
	numSet := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := numSet[num]; ok {
			return num
		}

		numSet[num] = struct{}{}
	}

	return 0
}

// Brute Force
// Time: O(N^2), Space: O(1)
func sln_1(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return nums[i]
			}
		}
	}
	return 0
}
