package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAtoi(t *testing.T) {
	var testCases = []struct {
		i string // input
		o int    // expected result
	}{
		//{
		//	"words and 987",
		//	0,
		//},
		//{
		//	"42",
		//	42,
		//},
		{
			" -42",
			-42,
		},
		//{
		//	"+-12",
		//	0,
		//},
	}

	for _, tt := range testCases {
		res := myAtoiReview(tt.i)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
