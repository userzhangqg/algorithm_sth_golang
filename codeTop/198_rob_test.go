package codeTop

import (
	"fmt"
	"reflect"
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
		}, {
			[]int{2, 7, 9, 3, 1},
			12,
		},
	}

	for _, tt := range testCases {
		res := robReview(tt.i)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
