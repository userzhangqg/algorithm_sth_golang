package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	var testCases = []struct {
		i []int   // input
		o [][]int // expected result
	}{
		{
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
	}

	for _, tt := range testCases {
		res := threeSum(tt.i)
		fmt.Println(res)
		//if res != tt.o {
		//	t.Errorf("error")
		//}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
