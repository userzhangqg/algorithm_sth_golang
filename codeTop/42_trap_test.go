package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTrap(t *testing.T) {
	var testCases = []struct {
		i []int // input
		o int   // expected result
	}{
		{
			[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			6,
		}, {
			[]int{4, 2, 0, 3, 2, 5},
			9,
		},
	}

	for _, tt := range testCases {
		res := trap(tt.i)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
