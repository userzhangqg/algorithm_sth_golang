package codeTop

import (
	"fmt"
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	var testCases = []struct {
		i []int // input
		o int   // expected result
	}{
		{
			[]int{1, 3, 6, 7, 9, 4, 10, 5, 6},
			6,
		}, {
			[]int{10, 9, 2, 5, 3, 7, 101, 18},
			4,
		}, {
			[]int{0, 1, 0, 3, 2, 3},
			4,
		}, {
			[]int{7, 7, 7, 7, 7, 7, 7},
			1,
		},
	}

	for _, tt := range testCases {
		res := lengthOfLIS(tt.i)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		//if !reflect.DeepEqual(res, tt.o) {
		//	t.Errorf("error")
		//}
	}
}
