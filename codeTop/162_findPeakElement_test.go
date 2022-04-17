package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindPeakElement(t *testing.T) {
	var testCases = []struct {
		i []int // input
		o int   // expected result
	}{
		{
			[]int{1, 2, 3, 1},
			2,
		}, {
			[]int{1, 2, 1, 3, 5, 6, 4},
			5,
		},
	}

	for _, tt := range testCases {
		res := findPeakElement(tt.i)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
