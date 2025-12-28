package p0074_search_a_2d_matrix

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	// Get Flat Size
	flatSize := m * n

	left := 0
	right := flatSize - 1

	for left <= right {
		mid := (left + right) / 2

		if target > matrix[mid/n][mid%n] {
			left = mid + 1
		} else if target < matrix[mid/n][mid%n] {
			right = mid - 1
		} else {
			return true
		}
	}

	return false
}
