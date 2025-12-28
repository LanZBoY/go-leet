package p0875_koko_eating_bananas

import (
	"testing"
)

func TestMinEatingSpeed(t *testing.T) {
	tests := []struct {
		name     string
		piles    []int
		h        int
		expected int
	}{
		{
			name:     "Example 1",
			piles:    []int{3, 6, 7, 11},
			h:        8,
			expected: 4,
		},
		{
			name:     "Example 2",
			piles:    []int{30, 11, 23, 4, 20},
			h:        5,
			expected: 30,
		},
		{
			name:     "Example 3",
			piles:    []int{30, 11, 23, 4, 20},
			h:        6,
			expected: 23,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := minEatingSpeed(tt.piles, tt.h)
			if actual != tt.expected {
				t.Errorf("minEatingSpeed(%v, %d) = %d; expected %d", tt.piles, tt.h, actual, tt.expected)
			}
		})
	}
}
