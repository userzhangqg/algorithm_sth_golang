package day06

import (
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	var testCases = []struct {
		nums    string // input
		newNums int    // expected result
	}{
		{
			"dvdf",
			3,
		},
		{
			"abcabcbb",
			3,
		},
		{
			"bbbbb",
			1,
		},
		{
			"pwwkew",
			3,
		},
		{
			"",
			0,
		},
		{
			" ",
			1,
		},
		{
			"au",
			2,
		},
	}

	for _, tt := range testCases {
		res := lengthOfLongestSubstring(tt.nums)
		fmt.Println(res)
		if !reflect.DeepEqual(res, tt.newNums) {
			t.Errorf("error")
		}
	}
}
