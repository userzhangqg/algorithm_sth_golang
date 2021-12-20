package day07

import (
	"fmt"
	"testing"
)

func TestArea(t *testing.T) {
	var testCases = []struct {
		nums    [][]int // input
		newNums int     // expected result
	}{
		{
			[][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}},
			6,
		},
	}

	for _, tt := range testCases {
		area := maxAreaOfIsland(tt.nums)
		fmt.Println(area)
		if area != tt.newNums {
			t.Errorf("error")
		}
	}
}
