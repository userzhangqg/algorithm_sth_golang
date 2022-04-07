package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHasPathSum(t *testing.T) {
	var testCases = []struct {
		i  *TreeNode // input
		i2 int
		o  bool // expected result
	}{
		{
			&TreeNode{1, &TreeNode{2, nil, nil}, nil},
			1,
			false,
		},
	}

	for _, tt := range testCases {
		res := hasPathSum(tt.i, tt.i2)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
