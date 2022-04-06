package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCoinChange(t *testing.T) {
	var testCases = []struct {
		i  []int // input
		i2 int
		o  int // expected result
	}{
		{
			[]int{2, 5, 10, 1},
			27,
			4,
		}, {
			[]int{1, 2, 5},
			11,
			3,
		}, {
			[]int{2},
			3,
			-1,
		}, {
			[]int{1},
			0,
			0,
		},
	}

	for _, tt := range testCases {
		res := coinChangeReview(tt.i, tt.i2)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
