package day03

import (
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	var testCases = []struct {
		nums    []int // input
		newNums []int // expected result
	}{
		{
			[]int{0, 1, 0, 3, 12},
			[]int{1, 3, 12, 0, 0},
		},
	}

	for _, tt := range testCases {
		moveZeroesDemo(tt.nums)
		fmt.Println(tt.nums)
		if !reflect.DeepEqual(tt.nums, tt.newNums) {
			t.Errorf("error")
		}
	}
}
