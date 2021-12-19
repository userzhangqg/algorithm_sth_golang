package day03

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	var testCases = []struct {
		nums    []int // input
		target  int
		newNums []int // expected result
	}{
		{
			[]int{2, 7, 11, 15},
			9,
			[]int{1, 2},
		},
		{
			[]int{2, 3, 4},
			6,
			[]int{1, 3},
		},
		{
			[]int{-1, 0},
			-1,
			[]int{1, 2},
		},
	}

	for _, tt := range testCases {
		result := twoSum(tt.nums, tt.target)
		fmt.Println(result)
		if !reflect.DeepEqual(result, tt.newNums) {
			t.Errorf("error")
		}
	}
}
