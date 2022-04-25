package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLargestNumber(t *testing.T) {
	var testCases = []struct {
		i []int  // input
		o string // expected result
	}{
		{
			[]int{10, 2},
			"210",
		}, {
			[]int{3, 30, 34, 5, 9},
			"9534330",
		},
	}

	for _, tt := range testCases {
		res := largestNumberReview(tt.i)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
