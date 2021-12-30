package day13

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsPowerOfTwo(t *testing.T) {
	var testCases = []struct {
		n   int  // input
		res bool // expected result
	}{
		{
			1,
			true,
		},
		{
			16,
			true,
		},
		{
			3,
			false,
		},
		{
			4,
			true,
		},
		{
			5,
			false,
		},
	}

	for _, tt := range testCases {
		res := isPowerOfTwo(tt.n)
		fmt.Println(res)
		if !reflect.DeepEqual(tt.res, res) {
			t.Errorf("error")
		}
	}
}
