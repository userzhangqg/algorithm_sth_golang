package day03

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMoveZeros(t *testing.T) {
	var fibTests = []struct {
		nums    []int // input
		newNums []int // expected result
	}{
		{
			[]int{0, 1, 0, 3, 12},
			[]int{1, 3, 12, 0, 0},
		},
		{
			[]int{0},
			[]int{0},
		},
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{0, 1},
			[]int{1, 0},
		},
		{
			[]int{0, 0, 0, 1, 0},
			[]int{1, 0, 0, 0, 0},
		},
		{
			[]int{1, 0, 1},
			[]int{1, 1, 0},
		},
		//{2, 1},
		//{3, 2},
		//{4, 3},
		//{5, 5},
		//{6, 8},
		//{7, 13},
	}

	for _, tt := range fibTests {
		//moveZeroes(tt.nums)
		moveZeroesDemo(tt.nums)
		fmt.Println(tt.nums)
		if !reflect.DeepEqual(tt.nums, tt.newNums) {
			t.Errorf("error")
		}
	}
}
