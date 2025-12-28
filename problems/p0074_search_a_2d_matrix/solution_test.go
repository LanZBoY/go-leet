package p0074_search_a_2d_matrix

import (
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		target int
		want   bool
	}{
		{
			name:   "Example 1",
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 3,
			want:   true,
		},
		{
			name:   "Example 2",
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 13,
			want:   false,
		},
		{
			name:   "Target smaller than smallest",
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 0,
			want:   false,
		},
		{
			name:   "Target larger than largest",
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 61,
			want:   false,
		},
		{
			name:   "Single element matrix - Found",
			matrix: [][]int{{1}},
			target: 1,
			want:   true,
		},
		{
			name:   "Single element matrix - Not Found",
			matrix: [][]int{{1}},
			target: 0,
			want:   false,
		},
		{
			name:   "Single row matrix - Found",
			matrix: [][]int{{1, 3, 5}},
			target: 3,
			want:   true,
		},
		{
			name:   "Single column matrix - Found",
			matrix: [][]int{{1}, {3}, {5}},
			target: 5,
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.matrix, tt.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
