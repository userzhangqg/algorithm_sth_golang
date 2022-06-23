package codeTopReview

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test",
			args{
				[]int{-1, 0, 1, 2, -1, -4},
			},
			[][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		// {
		//	"test",
		//	args{
		//		[]int{},
		//	},
		//	[][]int{},
		//}, {
		//	"test",
		//	args{
		//		[]int{0},
		//	},
		//	[][]int{},
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
