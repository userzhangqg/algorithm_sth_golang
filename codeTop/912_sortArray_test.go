package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSortArray(t *testing.T) {
	var testCases = []struct {
		i []int // input
		o []int // expected result
	}{
		{
			[]int{5, 2, 3, 1},
			[]int{1, 2, 3, 5},
		}, {
			[]int{5, 1, 1, 2, 0, 0},
			[]int{0, 0, 1, 1, 2, 5},
		},
	}

	for _, tt := range testCases {
		res := sortArray(tt.i)
		fmt.Println(res)
		//if res != tt.o {
		//	t.Errorf("error")
		//}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
