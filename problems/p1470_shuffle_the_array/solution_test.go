package p1470_shuffle_the_array

import (
	"reflect"
	"testing"
)

func TestShuffle(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		n    int
		want []int
	}{
		{
			name: "Example 1",
			nums: []int{2, 5, 1, 3, 4, 7},
			n:    3,
			want: []int{2, 3, 5, 4, 1, 7},
		},
		{
			name: "Example 2",
			nums: []int{1, 2, 3, 4, 4, 3, 2, 1},
			n:    4,
			want: []int{1, 4, 2, 3, 3, 2, 4, 1},
		},
		{
			name: "Example 3",
			nums: []int{1, 1, 2, 2},
			n:    2,
			want: []int{1, 2, 1, 2},
		},
		{
			name: "Minimum size n=1",
			nums: []int{1, 2},
			n:    1,
			want: []int{1, 2},
		},
		{
			name: "All same values",
			nums: []int{5, 5, 5, 5},
			n:    2,
			want: []int{5, 5, 5, 5},
		},
		{
			name: "Increasing x and y",
			nums: []int{1, 2, 3, 4, 5, 6},
			n:    3,
			want: []int{1, 4, 2, 5, 3, 6},
		},
		{
			name: "Decreasing x and y",
			nums: []int{6, 5, 4, 3, 2, 1},
			n:    3,
			want: []int{6, 3, 5, 2, 4, 1},
		},
		{
			name: "Maximum element value",
			nums: []int{100, 1, 99, 2},
			n:    2,
			want: []int{100, 99, 1, 2},
		},
		{
			name: "Alternating pattern",
			nums: []int{1, 3, 2, 4, 5, 7, 6, 8},
			n:    4,
			want: []int{1, 5, 3, 7, 2, 6, 4, 8},
		},
	}

	solutions := []struct {
		name string
		fn   func([]int, int) []int
	}{
		{"sln_1 (Direct indexing)", sln_1},
		{"sln_2 (Append pairs)", sln_2},
		{"sln_3 (Sequential pointer)", sln_3},
		{"sln_4 (Two-pass fill)", sln_4},
	}

	for _, sln := range solutions {
		t.Run(sln.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := sln.fn(tt.nums, tt.n); !reflect.DeepEqual(got, tt.want) {
						t.Errorf("%s: shuffle(%v, %d) = %v, want %v", sln.name, tt.nums, tt.n, got, tt.want)
					}
				})
			}
		})
	}

	t.Run("Main Function", func(t *testing.T) {
		for _, tt := range tests {
			if got := shuffle(tt.nums, tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shuffle(%v, %d) = %v, want %v", tt.nums, tt.n, got, tt.want)
			}
		}
	})
}
