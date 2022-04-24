package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxProduct(t *testing.T) {
	var testCases = []struct {
		i []int // input
		o int   // expected result
	}{
		{
			[]int{2, 3, -2, 4},
			6,
		}, {
			[]int{-2, 0, -1},
			0,
		},
	}

	for _, tt := range testCases {
		res := maxProductReview(tt.i)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
