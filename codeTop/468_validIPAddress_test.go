package codeTop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestValidIPAddress(t *testing.T) {
	var testCases = []struct {
		i string // input
		o string // expected result
	}{
		{
			"172.16.254.1",
			"IPv4",
		}, {
			"2001:0db8:85a3:0:0:8A2E:0370:7334",
			"IPv6",
		}, {
			"256.256.256.256",
			"Neither",
		},
	}

	for _, tt := range testCases {
		res := validIPAddressReview(tt.i)
		fmt.Println(res)
		if res != tt.o {
			t.Errorf("error")
		}
		if !reflect.DeepEqual(res, tt.o) {
			t.Errorf("error")
		}
	}
}
