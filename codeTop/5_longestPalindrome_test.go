package codeTop

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	var testCases = []struct {
		i string // input
		o string // expected result
	}{
		{
			"babad",
			"bab",
		}, {
			"cbba",
			"bb",
		},
	}

	for _, tt := range testCases {
		res := longestPalindrome(tt.i)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		//if !reflect.DeepEqual(res, tt.o) {
		//	t.Errorf("error")
		//}
	}
}
