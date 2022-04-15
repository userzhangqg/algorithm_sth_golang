package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaximalSquare(t *testing.T) {
	var testCases = []struct {
		i [][]byte // input
		o int      // expected result
	}{
		{
			[][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			4,
		}, {
			[][]byte{
				{'0', '1'},
				{'1', '0'},
			},
			1,
		}, {
			[][]byte{
				{'0'},
			},
			0,
		},
	}

	for _, tt := range testCases {
		res := maximalSquare(tt.i)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
