package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMerge56(t *testing.T) {
	var testCases = []struct {
		i [][]int // input
		o [][]int // expected result
	}{
		{
			[][]int{
				{1, 3},
				{2, 6},
				{8, 10},
				{15, 18},
			},
			[][]int{
				{1, 6},
				{8, 10},
				{15, 18},
			},
		}, {
			[][]int{
				{1, 4},
				{4, 5},
			},
			[][]int{
				{1, 5},
			},
		},
	}

	for _, tt := range testCases {
		res := merge56(tt.i)
		fmt.Println(res)
		//if res != tt.o {
		//	t.Errorf("error")
		//}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
