package p1929_concatenation_of_array

func getConcatenation(nums []int) []int {
	return sln_1(nums)
}

// 已知結果長度時，實務上建議先 make 配好記憶體再填入，而非一路 append 擴容。
// append 在 cap 不足時會重新分配並 copy 舊資料；大 slice 或高頻路徑下，
// 多次分配的成本（含 allocator 壓力，必要時才會走到 mmap 等較重的 kernel 路徑）
// 會比一次 make + 直接寫入來得明顯。
// sln_1 / sln_2 / sln_4 採預先分配；sln_3 最簡短但可能觸發擴容，不適合當 production 預設。

// Modulo indexing
// Time: O(n), Space: O(n)
// 預先 make(2*n)，單次分配、無擴容。
func sln_1(nums []int) []int {
	n := len(nums)
	ans := make([]int, 2*n)

	for i := 0; i < 2*n; i++ {
		ans[i] = nums[i%n]
	}

	return ans
}

// copy twice
// Time: O(n), Space: O(n)
// 預先 make(2*n) + copy；語意清楚，實務上最常推薦的寫法之一。
func sln_2(nums []int) []int {
	n := len(nums)
	ans := make([]int, 2*n)

	copy(ans[:n], nums)
	copy(ans[n:], nums)

	return ans
}

// append
// Time: O(n), Space: O(n)
// 一行搞定，但依 nums 的 cap 可能觸發 0～多次擴容與 copy；已知長度時不建議。
func sln_3(nums []int) []int {
	return append(nums, nums...)
}

// two explicit loops
// Time: O(n), Space: O(n)
// 預先 make(2*n)，兩段迴圈分別填前半與後半。
func sln_4(nums []int) []int {
	n := len(nums)
	ans := make([]int, 2*n)

	for i := 0; i < n; i++ {
		ans[i] = nums[i]
	}
	for i := 0; i < n; i++ {
		ans[i+n] = nums[i]
	}

	return ans
}
