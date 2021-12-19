package day04

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverseWords(t *testing.T) {
	var testCases = []struct {
		nums    string // input
		newNums string // expected result
	}{
		{
			"Let's take LeetCode contest",
			"s'teL ekat edoCteeL tsetnoc",
		},
	}

	for _, tt := range testCases {
		newString := reverseWords(tt.nums)
		fmt.Println(newString)
		if !reflect.DeepEqual(newString, tt.newNums) {
			t.Errorf("error")
		}
	}
}
