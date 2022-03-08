package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	var testCases = []struct {
		i  []int // input
		i2 int
		o  []int // expected result
	}{
		{
			[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		},
	}

	for _, tt := range testCases {
		res := twoSum(tt.i, tt.i2)
		fmt.Println(res)
		//if res != tt.o {
		//	t.Errorf("error")
		//}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
