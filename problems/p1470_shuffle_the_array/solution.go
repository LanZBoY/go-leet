package p1470_shuffle_the_array

func shuffle(nums []int, n int) []int {
	return sln_1(nums, n)
}

// sln_1 / sln_3 / sln_4 預先 make(2*n) 再填入；已知長度時比 append 擴容更穩定。

// Direct indexing (i*2)
// Time: O(n), Space: O(n)
// 每輪 i 直接定位到 ans 的 2i、2i+1，最精簡。
func sln_1(nums []int, n int) []int {
	ans := make([]int, 2*n)

	for i := 0; i < n; i++ {
		ans[i*2] = nums[i]
		ans[i*2+1] = nums[i+n]
	}

	return ans
}

// Append pairs
// Time: O(n), Space: O(n)
// 語意最直覺：每輪 append 一對 (x, y)。cap 預配好則不擴容。
func sln_2(nums []int, n int) []int {
	ans := make([]int, 0, 2*n)

	for i := 0; i < n; i++ {
		ans = append(ans, nums[i], nums[i+n])
	}

	return ans
}

// Sequential write pointer
// Time: O(n), Space: O(n)
// 用 j 從左到右依序填，不用背 i*2 公式。
func sln_3(nums []int, n int) []int {
	ans := make([]int, 2*n)
	j := 0

	for i := 0; i < n; i++ {
		ans[j] = nums[i]
		j++
		ans[j] = nums[i+n]
		j++
	}

	return ans
}

// Two-pass fill
// Time: O(n), Space: O(n)
// 先填所有偶數位 (x)，再填所有奇數位 (y)。
func sln_4(nums []int, n int) []int {
	ans := make([]int, 2*n)

	for i := 0; i < n; i++ {
		ans[i*2] = nums[i]
	}
	for i := 0; i < n; i++ {
		ans[i*2+1] = nums[i+n]
	}

	return ans
}
