package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	var testCases = []struct {
		i  []int // input
		i2 int
		o  [][]int // expected result
	}{
		{
			[]int{2, 3, 6, 7},
			7,
			[][]int{
				{
					2, 2, 3,
				},
				{
					7,
				},
			},
		},
	}

	for _, tt := range testCases {
		res := combinationSum(tt.i, tt.i2)
		fmt.Println(res)
		//if res != tt.o {
		//	t.Errorf("error")
		//}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
