package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMultiply(t *testing.T) {
	var testCases = []struct {
		i  string // input
		i2 string // input
		o  string // expected result
	}{
		{
			"9",
			"9",
			"81",
		}, {
			"2",
			"3",
			"6",
		}, {
			"123",
			"456",
			"56088",
		},
	}

	for _, tt := range testCases {
		res := multiplyReview(tt.i, tt.i2)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
