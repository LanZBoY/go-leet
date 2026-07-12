package p1929_concatenation_of_array

import (
	"reflect"
	"testing"
)

func TestGetConcatenation(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 1},
			want: []int{1, 2, 1, 1, 2, 1},
		},
		{
			name: "Example 2",
			nums: []int{1, 3, 2, 1},
			want: []int{1, 3, 2, 1, 1, 3, 2, 1},
		},
		{
			name: "Single element",
			nums: []int{5},
			want: []int{5, 5},
		},
		{
			name: "Two elements",
			nums: []int{1, 2},
			want: []int{1, 2, 1, 2},
		},
		{
			name: "All same values",
			nums: []int{7, 7, 7},
			want: []int{7, 7, 7, 7, 7, 7},
		},
		{
			name: "Increasing sequence",
			nums: []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5},
		},
		{
			name: "Maximum element value",
			nums: []int{1000, 999, 1000},
			want: []int{1000, 999, 1000, 1000, 999, 1000},
		},
	}

	solutions := []struct {
		name string
		fn   func([]int) []int
	}{
		{"sln_1 (Modulo indexing)", sln_1},
		{"sln_2 (copy)", sln_2},
		{"sln_3 (append)", sln_3},
		{"sln_4 (two loops)", sln_4},
	}

	for _, sln := range solutions {
		t.Run(sln.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := sln.fn(tt.nums); !reflect.DeepEqual(got, tt.want) {
						t.Errorf("%s: getConcatenation(%v) = %v, want %v", sln.name, tt.nums, got, tt.want)
					}
				})
			}
		})
	}

	t.Run("Main Function", func(t *testing.T) {
		for _, tt := range tests {
			if got := getConcatenation(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getConcatenation(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		}
	})
}
