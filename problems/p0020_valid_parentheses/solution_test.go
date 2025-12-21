package p0020_valid_parentheses

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "Example 1",
			s:    "()",
			want: true,
		},
		{
			name: "Example 2",
			s:    "()[]{}",
			want: true,
		},
		{
			name: "Example 3",
			s:    "(]",
			want: false,
		},
		{
			name: "Example 4",
			s:    "([])",
			want: true,
		},
		{
			name: "Example 5",
			s:    "([)]",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
