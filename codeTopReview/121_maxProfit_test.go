package codeTopReview

import "testing"

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
			"test",
			args{
				[]int{7, 1, 5, 3, 6, 4},
			},
			5,
		}, {
			"test",
			args{
				[]int{7, 6, 4, 3, 1},
			},
			0,
		}, {
			"test",
			args{
				[]int{7, 2, 5, 1, 3, 6, 4},
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
