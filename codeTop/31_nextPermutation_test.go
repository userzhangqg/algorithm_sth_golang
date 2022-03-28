package codeTop

import (
	"fmt"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"",
			args{
				[]int{1, 2, 3}, // got 1,3,2
			},
		}, {
			"",
			args{
				[]int{3, 2, 1}, // got 1,2,3
			},
		}, {
			"",
			args{
				[]int{1, 1, 5}, // got 1,5,1
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.args.nums)
			fmt.Println(tt.args.nums)
		})
	}
}
