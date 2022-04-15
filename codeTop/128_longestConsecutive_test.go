package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	var testCases = []struct {
		i []int // input
		o int   // expected result
	}{
		{
			[]int{100, 4, 200, 1, 3, 2},
			4,
		}, {
			[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			9,
		},
	}

	for _, tt := range testCases {
		res := longestConsecutive(tt.i)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
