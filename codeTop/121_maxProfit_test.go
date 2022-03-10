package codeTop

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	var cases = []struct {
		i []int
		o int
	}{
		{
			[]int{7, 1, 5, 3, 6, 4},
			5,
		},
		{
			[]int{7, 6, 4, 3, 1},
			0,
		},
		{
			[]int{7, 2, 5, 1, 4, 5},
			4,
		},
	}

	for _, tt := range cases {
		ans := maxProfit(tt.i)
		fmt.Println(ans)

		if ans != tt.o {
			t.Errorf("error")
		}
	}
}
