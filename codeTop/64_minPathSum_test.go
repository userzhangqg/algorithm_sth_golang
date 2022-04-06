package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMinPathSum(t *testing.T) {
	var testCases = []struct {
		i [][]int // input
		o int     // expected result
	}{
		{
			[][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			},
			7,
		}, {
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			12,
		},
	}

	for _, tt := range testCases {
		res := minPathSumReview(tt.i)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
