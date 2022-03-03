package codeTop

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	var testCases = []struct {
		i string // input
		o int    // expected result
	}{
		{
			"abcabcbb",
			3,
		}, {
			"bbbbb",
			1,
		}, {
			"pwwkew",
			3,
		},
	}

	for _, tt := range testCases {
		res := lengthOfLongestSubstring(tt.i)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		//if !reflect.DeepEqual(res, tt.o) {
		//	t.Errorf("error")
		//}
	}
}
