package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindMin(t *testing.T) {
	var testCases = []struct {
		i []int // input
		o int   // expected result
	}{
		{
			[]int{5, 1, 2, 3, 4},
			1,
		}, {
			[]int{3, 4, 5, 1, 2},
			1,
		}, {
			[]int{4, 5, 6, 7, 0, 1, 2},
			0,
		}, {
			[]int{11, 13, 15, 17},
			11,
		},
	}

	for _, tt := range testCases {
		res := findMin(tt.i)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
