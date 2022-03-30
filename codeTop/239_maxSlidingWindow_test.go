package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	var testCases = []struct {
		i []int // input
		k int
		o []int // expected result
	}{
		{
			[]int{1, 3, -1, -3, 5, 3, 6, 7},
			3,
			[]int{3, 3, 5, 5, 6, 7},
		}, {
			[]int{1},
			1,
			[]int{1},
		},
	}

	for _, tt := range testCases {
		res := maxSlidingWindow(tt.i, tt.k)
		fmt.Println(res)
		//if res != tt.o {
		//	t.Errorf("error")
		//}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
