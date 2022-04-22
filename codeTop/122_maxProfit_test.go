package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxProfit122(t *testing.T) {
	var testCases = []struct {
		i []int // input
		o int   // expected result
	}{
		//{
		//	[]int{2,1,2,0,1},
		//	2,
		//},
		{
			[]int{7, 1, 5, 3, 6, 4},
			7,
		}, {
			[]int{1, 2, 3, 4, 5},
			4,
		}, {
			[]int{7, 6, 4, 3, 1},
			0,
		},
	}

	for _, tt := range testCases {
		res := maxProfitReview(tt.i)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
