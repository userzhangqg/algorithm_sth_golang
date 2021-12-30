package day12

import (
	"fmt"
	"testing"
)

func TestRob(t *testing.T) {
	var testCases = []struct {
		i []int // input
		o int   // expected result
	}{
		{
			[]int{1, 2, 3, 1},
			4,
		},
		{
			[]int{2, 7, 9, 3, 1},
			12,
		},
		{
			[]int{2, 1, 1, 2},
			4,
		},
	}

	for _, tt := range testCases {
		res := rob(tt.i)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
	}
}
