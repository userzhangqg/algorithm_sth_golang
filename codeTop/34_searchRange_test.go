package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	var testCases = []struct {
		i  []int // input
		i2 int
		o  []int // expected result
	}{
		{
			[]int{5, 7, 7, 8, 8, 19},
			8,
			[]int{3, 4},
		},
	}

	for _, tt := range testCases {
		res := searchRange(tt.i, tt.i2)
		fmt.Println(res)
		//if res != tt.o {
		//	t.Errorf("error")
		//}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
