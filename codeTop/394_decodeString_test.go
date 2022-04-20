package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDecodeString(t *testing.T) {
	var testCases = []struct {
		i string // input
		o string // expected result
	}{
		{
			"3[z]2[2[y]pq4[2[jk]e1[f]]]ef",
			"zzzyypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkefef",
		}, {
			"100[leetcode]",
			"leetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcode",
		}, {
			"3[a]2[bc]",
			"aaabcbc",
		}, {
			"3[a2[c]]",
			"accaccacc",
		}, {
			"2[abc]3[cd]ef",
			"abcabccdcdcdef",
		}, {
			"abc3[cd]xyz",
			"abccdcdcdxyz",
		},
	}

	for _, tt := range testCases {
		res := decodeStringReview(tt.i)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
