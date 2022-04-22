package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCalculate(t *testing.T) {
	var testCases = []struct {
		i string // input
		o int    // expected result
	}{
		{
			"3+2*2",
			7,
		}, {
			" 3/2 ",
			1,
		}, {
			" 3+5 / 2 ",
			5,
		},
	}

	for _, tt := range testCases {
		res := calculateReview(tt.i)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
