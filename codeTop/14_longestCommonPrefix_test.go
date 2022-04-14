package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	var testCases = []struct {
		i []string // input
		o string   // expected result
	}{
		{
			[]string{"a"},
			"a",
		}, {
			[]string{"dog", "racecar", "car"},
			"",
		}, {
			[]string{"flower", "flow", "flight"},
			"fl",
		},
	}

	for _, tt := range testCases {
		res := longestCommonPrefixReview(tt.i)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
