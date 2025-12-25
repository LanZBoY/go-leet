package p0704_binary_search

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "Example 1",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			want:   4,
		},
		{
			name:   "Example 2",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 2,
			want:   -1,
		},
		{
			name:   "Single element found",
			nums:   []int{5},
			target: 5,
			want:   0,
		},
		{
			name:   "Single element not found",
			nums:   []int{5},
			target: 2,
			want:   -1,
		},
		{
			name:   "Two elements, find first",
			nums:   []int{2, 5},
			target: 2,
			want:   0,
		},
		{
			name:   "Two elements, find second",
			nums:   []int{2, 5},
			target: 5,
			want:   1,
		},
		{
			name:   "Two elements, not found",
			nums:   []int{2, 5},
			target: 3,
			want:   -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.nums, tt.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
