package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
	var testCases = []struct {
		i []int   // input
		o [][]int // expected result
	}{
		{
			[]int{1, 2, 3},
			[][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		}, {
			[]int{0, 1},
			[][]int{
				{0, 1},
				{1, 0},
			},
		}, {
			[]int{1},
			[][]int{
				{1},
			},
		},
	}

	for _, tt := range testCases {
		ans = [][]int{}
		//res := permute(tt.i)
		res := permuteReview(tt.i)
		fmt.Println(res)
		//if res != tt.o {
		//	t.Errorf("error")
		//}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
