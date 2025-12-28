package p0003_longest_substring_without_repeating_characters

import (
	"reflect"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{s: "abcabcbb"},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{s: "bbbbb"},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{s: "pwwkew"},
			want: 3,
		},
		{
			name: "Empty string",
			args: args{s: ""},
			want: 0,
		},
		{
			name: "Space character",
			args: args{s: " "},
			want: 1,
		},
		{
			name: "All unique",
			args: args{s: "abcdef"},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
