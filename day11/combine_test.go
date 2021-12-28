package day11

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var testCases = []struct {
		n int // input
		k int // expected result
	}{
		{
			4,
			2,
		},
	}

	for _, tt := range testCases {
		res := combine(tt.n, tt.k)
		fmt.Println(res)
		//if !reflect.DeepEqual(tt.nums, tt.newNums) {
		//	t.Errorf("error")
		//}
	}
}
