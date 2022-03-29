package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFirstMissingPositive(t *testing.T) {
	var testCases = []struct {
		i []int // input
		o int   // expected result
	}{
		{
			[]int{1, 2, 0},
			3,
		}, {
			[]int{3, 4, -1, 1},
			2,
		}, {
			[]int{7, 8, 9, 11, 12},
			1,
		},
	}

	for _, tt := range testCases {
		res := firstMissingPositive(tt.i)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
