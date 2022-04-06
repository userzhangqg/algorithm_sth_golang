package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
	var testCases = []struct {
		i []int   // input
		o [][]int // expected result
	}{
		{
			[]int{1, 2, 3},
			[][]int{
				{},
				{1},
				{2},
				{1, 2},
				{3},
				{1, 3},
				{2, 3},
				{1, 2, 3},
			},
		}, {
			[]int{0},
			[][]int{
				{},
				{0},
			},
		},
	}

	for _, tt := range testCases {
		res := subsetsReview(tt.i)
		fmt.Println(res)
		//if res != tt.o {
		//	t.Errorf("error")
		//}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
