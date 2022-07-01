package codeTopReview

import (
	"fmt"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"test",
			args{
				[]int{1, 2, 3, 0, 0, 0}, 3,
				[]int{2, 5, 6}, 3,
			},
		}, {
			"test",
			args{
				[]int{1}, 1,
				[]int{}, 0,
			},
		}, {
			"test",
			args{
				[]int{0}, 0,
				[]int{1}, 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.args.nums1)
		})
	}
}
