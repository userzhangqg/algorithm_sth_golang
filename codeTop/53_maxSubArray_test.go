package codeTop

import (
	"fmt"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	var testCases = []struct {
		i []int // input
		o int   // expected result
	}{
		{
			[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			6,
		}, {
			[]int{1},
			1,
		}, {
			[]int{5, 4, -1, 7, 8},
			23,
		},
	}

	for _, tt := range testCases {
		res := maxSubArray(tt.i)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		//if !reflect.DeepEqual(res, tt.o) {
		//	t.Errorf("error")
		//}
	}
}
