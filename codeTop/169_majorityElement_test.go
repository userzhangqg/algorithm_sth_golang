package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	var testCases = []struct {
		i []int // input
		o int   // expected result
	}{
		{
			[]int{3, 2, 3},
			3,
		}, {
			[]int{2, 2, 1, 1, 1, 2, 2},
			2,
		}, {
			[]int{3, 2, 2, 3, 1, 1, 3, 3, 3},
			3,
		},
	}

	for _, tt := range testCases {
		res := majorityElement(tt.i)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
