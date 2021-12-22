package day09

import (
	"fmt"
	"reflect"
	"testing"
)

func TestOranges(t *testing.T) {
	var testCases = []struct {
		nums    [][]int // input
		newNums int     // expected result
	}{
		{
			[][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}},
			4,
		}, {
			[][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}},
			-1,
		}, {
			[][]int{{0, 2}},
			0,
		},
	}

	for _, tt := range testCases {
		res := orangesRotting(tt.nums)
		fmt.Println(res)
		if !reflect.DeepEqual(res, tt.newNums) {
			t.Errorf("error")
		}
	}
}
