package day13

import (
	"fmt"
	"testing"
)

func TestHamming(t *testing.T) {
	var testCases = []struct {
		num uint32 // input
		res int    // expected result
	}{
		{
			0b00000000000000000000000000001011,
			3,
		},
		{
			0b00000000000000000000000010000000,
			1,
		},
		{
			0b11111111111111111111111111111101,
			31,
		},
	}

	for _, tt := range testCases {
		res := hammingWeight(tt.num)
		fmt.Println(res)
		if res != tt.res {
			t.Errorf("error")
		}
	}
}
