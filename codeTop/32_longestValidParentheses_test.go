package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLongestValidParentheses(t *testing.T) {
	var testCases = []struct {
		i string // input
		o int    // expected result
	}{
		{
			"()()",
			4,
		}, {
			"(()",
			2,
		},
		{
			"()(())",
			6,
		},
		{
			")()())",
			4,
		}, {
			"",
			0,
		}, {
			"()(()",
			2,
		},
	}

	for _, tt := range testCases {
		res := longestValidParenthesesReview(tt.i)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
