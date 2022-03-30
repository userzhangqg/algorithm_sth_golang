package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMinWindow(t *testing.T) {
	var testCases = []struct {
		s string // input
		t string // expected result
		o string
	}{
		{
			"ADOBECODEBANC",
			"ABC",
			"BANC",
		},
		{
			"a",
			"a",
			"a",
		}, {
			"a",
			"aa",
			"",
		},
	}

	for _, tt := range testCases {
		res := minWindow(tt.s, tt.t)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
