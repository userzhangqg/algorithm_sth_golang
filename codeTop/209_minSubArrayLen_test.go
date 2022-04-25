package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	var testCases = []struct {
		i1 int   // expected result
		i2 []int // input
		o  int
	}{
		{
			15,
			[]int{5, 1, 3, 5, 10, 7, 4, 9, 2, 8},
			2,
		}, {
			7,
			[]int{2, 3, 1, 2, 4, 3},
			2,
		}, {
			4,
			[]int{1, 4, 4},
			1,
		}, {
			11,
			[]int{1, 1, 1, 1, 1, 1, 1, 1},
			0,
		},
	}

	for _, tt := range testCases {
		res := minSubArrayLenReview(tt.i1, tt.i2)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
