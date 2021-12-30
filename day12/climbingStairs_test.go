package day12

import (
	"fmt"
	"testing"
)

func TestClimbing(t *testing.T) {
	var testCases = []struct {
		i int // input
		o int // expected result
	}{
		{
			2,
			2,
		},
		{
			3,
			3,
		},
		{
			4,
			5,
		},
		//{
		//	44,
		//	5,
		//},
	}

	for _, tt := range testCases {
		res := climbStairs(tt.i)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
	}
}
