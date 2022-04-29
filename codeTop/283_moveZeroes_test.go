package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	var testCases = []struct {
		i []int // input
		o []int // expected result
	}{
		{
			[]int{0, 1, 0, 3, 12},
			[]int{1, 3, 12, 0, 0},
		},
	}

	for _, tt := range testCases {
		moveZeroes(tt.i)
		fmt.Println(tt.i)
		//if tt.i != tt.o {
		//	t.Errorf("error")
		//}
		if !reflect.DeepEqual(tt.i, tt.o) {
			t.Errorf("error")
		}
	}
}
