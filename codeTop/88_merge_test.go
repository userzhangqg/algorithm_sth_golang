package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

/*
nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
*/

func TestMerge(t *testing.T) {
	cases := []struct {
		i1 []int
		s1 int
		i2 []int
		s2 int
		o  []int
	}{
		{
			[]int{1, 2, 3, 0, 0, 0},
			3,
			[]int{2, 5, 6},
			3,
			[]int{1, 2, 2, 3, 5, 6},
		}, {
			[]int{1},
			1,
			[]int{},
			0,
			[]int{1},
		}, {
			[]int{0},
			0,
			[]int{1},
			1,
			[]int{1},
		},
	}

	for _, tt := range cases {
		merge(tt.i1, tt.s1, tt.i2, tt.s2)
		fmt.Println(tt.i1)

		if !reflect.DeepEqual(tt.i1, tt.o) {
			t.Errorf("error")
		}
	}
}
