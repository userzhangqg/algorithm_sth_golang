package day11

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	var testCases = []struct {
		nums []int // input
		//newNums []int // expected result
	}{
		{
			[]int{1, 2, 3},
			//[]int{1, 3, 12, 0, 0},
		},
	}

	for _, tt := range testCases {
		res := permute(tt.nums)
		fmt.Println(res)
		//if !reflect.DeepEqual(tt.nums, tt.newNums) {
		//	t.Errorf("error")
		//}
	}
}
