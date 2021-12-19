package day06

import (
	"fmt"
	"testing"
)

func TestInclusion(t *testing.T) {
	var testCases = []struct {
		s1          string // input
		s2          string // expected result
		isInclusion bool
	}{
		{
			"abc",
			"bbbca",
			true,
		},
		{
			"ab",
			"eidbaooo",
			true,
		},
	}

	for _, tt := range testCases {
		isInclusion := checkInclusion(tt.s1, tt.s2)
		fmt.Println(isInclusion)
		if isInclusion != tt.isInclusion {
			t.Errorf("error")
		}
	}
}
