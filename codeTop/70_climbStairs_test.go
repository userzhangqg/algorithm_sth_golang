package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	var testCases = []struct {
		i int // input
		o int // expected result
	}{
		{
			2,
			2,
		}, {
			3,
			3,
		},
	}

	for _, tt := range testCases {
		res := climbStairs(tt.i)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
