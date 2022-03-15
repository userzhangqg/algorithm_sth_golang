package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddStrings(t *testing.T) {
	var testCases = []struct {
		nums1 string // input
		nums2 string // input
		o     string // expected result
	}{
		{
			"11",
			"123",
			"134",
		}, {
			"456",
			"77",
			"533",
		}, {
			"0",
			"0",
			"0",
		},
	}

	for _, tt := range testCases {
		res := addStrings(tt.nums1, tt.nums2)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
