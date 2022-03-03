package codeTop

import (
	"fmt"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	var testCases = []struct {
		i  []int // input
		i2 int
		o  int // expected result
	}{
		{
			[]int{3, 2, 1, 5, 6, 4},
			2,
			5,
		},
		{
			[]int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			4,
			4,
		},
	}

	for _, tt := range testCases {
		res := findKthLargest(tt.i, tt.i2)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		//if !reflect.DeepEqual(res, tt.o) {
		//	t.Errorf("error")
		//}
	}
}
