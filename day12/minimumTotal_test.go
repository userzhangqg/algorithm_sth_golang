package day12

import (
	"fmt"
	"testing"
)

func TestTriangle(t *testing.T) {
	var testCases = []struct {
		triangle [][]int // input
		res      int     // expected result
	}{
		{
			[][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}},
			11,
		},
		{
			[][]int{{-10}},
			-10,
		},
	}

	for _, tt := range testCases {
		res := minimumTotal(tt.triangle)
		fmt.Println(res)
		if res != tt.res {
			t.Errorf("error")
		}
	}
}
