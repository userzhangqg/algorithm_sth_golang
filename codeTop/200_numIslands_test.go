package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNumIslands(t *testing.T) {
	var testCases = []struct {
		i [][]byte // input
		o int      // expected result
	}{
		{
			[][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			1,
		},
	}

	for _, tt := range testCases {
		res := numIslands(tt.i)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
