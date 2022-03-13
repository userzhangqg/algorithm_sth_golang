package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	var testCases = []struct {
		i  []int // input
		i2 int
		o  int // expected result
	}{
		{
			[]int{4, 5, 6, 7, 0, 1, 2},
			0,
			4,
		}, {
			[]int{4, 5, 6, 7, 0, 1, 2},
			3,
			-1,
		}, {
			[]int{1},
			0,
			-1,
		},
	}

	for _, tt := range testCases {
		res := search33(tt.i, tt.i2)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
