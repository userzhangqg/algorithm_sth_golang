package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	var testCases = []struct {
		i  [][]int // input
		i2 int
		o  bool // expected result
	}{
		{
			[][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			4,
			true,
		},
	}

	for _, tt := range testCases {
		res := searchMatrix(tt.i, tt.i2)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
