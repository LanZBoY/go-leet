package p0121_best_time_to_buy_and_sell_stock

import (
	"reflect"
	"testing"
)

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{prices: []int{7, 1, 5, 3, 6, 4}},
			want: 5,
		},
		{
			name: "Example 2",
			args: args{prices: []int{7, 6, 4, 3, 1}},
			want: 0,
		},
		{
			name: "Empty prices",
			args: args{prices: []int{}},
			want: 0,
		},
		{
			name: "Single price",
			args: args{prices: []int{1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
