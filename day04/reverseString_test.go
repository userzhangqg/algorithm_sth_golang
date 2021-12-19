package day04

import (
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	var testCases = []struct {
		nums    []string // input
		newNums []string // expected result
	}{
		{
			[]string{"h", "e", "l", "l", "o"},
			[]string{"o", "l", "l", "e", "h"},
		},
		{
			[]string{"H", "e", "l", "l", "O"},
			[]string{"O", "l", "l", "e", "H"},
		},
		{
			[]string{"H", "o"},
			[]string{"o", "H"},
		},
	}

	for _, tt := range testCases {
		reverseString(tt.nums)
		fmt.Println(tt.nums)
		if !reflect.DeepEqual(tt.nums, tt.newNums) {
			t.Errorf("error")
		}
	}
}
