package p0287_find_the_duplicate_number

import (
	"testing"
)

func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 3, 4, 2, 2},
			expected: 2,
		},
		{
			name:     "Example 2",
			nums:     []int{3, 1, 3, 4, 2},
			expected: 3,
		},
		{
			name:     "Example 3",
			nums:     []int{3, 3, 3, 3, 3},
			expected: 3,
		},
		{
			name:     "Minimum size",
			nums:     []int{1, 1},
			expected: 1,
		},
		{
			name:     "Duplicate is 1 and n",
			nums:     []int{1, 2, 3, 1},
			expected: 1,
		},
		{
			name:     "Duplicate in the middle",
			nums:     []int{2, 1, 3, 2, 4},
			expected: 2,
		},
		{
			name:     "All elements same",
			nums:     []int{2, 2, 2, 2, 2},
			expected: 2,
		},
	}

	solutions := []struct {
		name string
		fn   func([]int) int
	}{
		{"sln_1 (Brute Force)", sln_1},
		{"sln_2 (Hash Map)", sln_2},
		{"sln_3 (Binary Search)", sln_3},
		{"sln_4 (Cycle Detection)", sln_4},
	}

	for _, sln := range solutions {
		t.Run(sln.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					actual := sln.fn(tt.nums)
					if actual != tt.expected {
						t.Errorf("%s: findDuplicate(%v) = %d; expected %d", sln.name, tt.nums, actual, tt.expected)
					}
				})
			}
		})
	}

	// Verify the main function points to the optimal solution (which should satisfy constraints best)
	// We can't easily check identity, but we just run cases on main too just in case.
	t.Run("Main Function", func(t *testing.T) {
		for _, tt := range tests {
			actual := findDuplicate(tt.nums)
			if actual != tt.expected {
				t.Errorf("findDuplicate(%v) = %d; expected %d", tt.nums, actual, tt.expected)
			}
		}
	})
}

func BenchmarkSolutions(b *testing.B) {
	nums := []int{1, 3, 4, 2, 2}
	b.Run("sln_1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sln_1(nums)
		}
	})
	b.Run("sln_2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sln_2(nums)
		}
	})
	b.Run("sln_3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sln_3(nums)
		}
	})
	b.Run("sln_4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sln_4(nums)
		}
	})
}
