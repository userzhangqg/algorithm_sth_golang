package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUniquePaths(t *testing.T) {
	var testCases = []struct {
		i  int // input
		i2 int
		o  int // expected result
	}{
		{
			3,
			7,
			28,
		}, {
			3,
			2,
			3,
		}, {
			7,
			3,
			28,
		}, {
			3,
			3,
			6,
		},
	}

	for _, tt := range testCases {
		res := uniquePaths(tt.i, tt.i2)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
