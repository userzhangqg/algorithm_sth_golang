package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCompareVersion(t *testing.T) {
	var testCases = []struct {
		i  string // input
		i2 string
		o  int // expected result
	}{
		{
			"1.0",
			"1.1",
			-1,
		}, {
			"0",
			"1",
			-1,
		}, {
			"1.0.1",
			"1",
			1,
		}, {
			"1.01",
			"1.001",
			0,
		}, {
			"1.0",
			"1.0.0",
			0,
		}, {
			"0.1",
			"1.1",
			-1,
		},
	}

	for _, tt := range testCases {
		res := compareVersion(tt.i, tt.i2)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
