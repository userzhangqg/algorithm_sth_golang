package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMySqrt(t *testing.T) {
	var testCases = []struct {
		i int // input
		o int // expected result
	}{
		{
			4,
			2,
		}, {
			8,
			2,
		},
	}

	for _, tt := range testCases {
		res := mySqrt(tt.i)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
