package codeTopReview

import "testing"

func Test_search704(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test",
			args{
				[]int{-1, 0, 3, 5, 9, 12}, 9,
			},
			4,
		}, {
			"test",
			args{
				[]int{-1, 0, 3, 5, 9, 12}, 2,
			},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search704(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search704() = %v, want %v", got, tt.want)
			}
		})
	}
}
