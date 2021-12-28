package day11

import (
	"fmt"
	"testing"
)

func TestLetter(t *testing.T) {
	var testCases = []struct {
		S string // input
		//newNums []int // expected result
	}{
		{
			"a1b2",
		},
		//{
		//	"3z4",
		//},
	}

	for _, tt := range testCases {
		res := letterCasePermutation(tt.S)
		//res := letterCasePermutationDemo(tt.S)
		fmt.Println(res)
		//if !reflect.DeepEqual(tt.nums, tt.newNums) {
		//	t.Errorf("error")
		//}
	}
}
