package codeTop

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	var testCases = []struct {
		i string // input
		o bool   // expected result
	}{
		{
			"()",
			true,
		}, {
			"()[]{}",
			true,
		}, {
			"(]",
			false,
		}, {
			"([)]",
			false,
		}, {
			"{[]}",
			true,
		},
	}

	for _, tt := range testCases {
		res := isValid(tt.i)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		//if !reflect.DeepEqual(res, tt.o) {
		//	t.Errorf("error")
		//}
	}
}
