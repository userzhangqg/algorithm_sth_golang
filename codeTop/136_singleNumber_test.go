package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	var testCases = []struct {
		i []int // input
		o int   // expected result
	}{
		{
			[]int{2, 2, 1},
			1,
		}, {
			[]int{4, 1, 2, 1, 2},
			4,
		},
	}

	for _, tt := range testCases {
		res := singleNumber(tt.i)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
