package day14

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var testCases = []struct {
		n   uint32 // input
		res uint32 // expected result
	}{
		{
			0b00000010100101000001111010011100,
			0b00111001011110000010100101000000,
		},
		{
			0b11111111111111111111111111111101,
			0b10111111111111111111111111111111,
		},
	}

	for _, tt := range testCases {
		res := reverseBits(tt.n)
		fmt.Println(res)
		if res != tt.res {
			t.Errorf("error")
		}
	}
}
