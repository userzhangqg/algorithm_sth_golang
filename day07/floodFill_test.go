package day07

import (
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	var testCases = []struct {
		image    [][]int // input
		sr       int     // expected result
		sc       int
		newColor int
		newImage [][]int
	}{
		{
			[][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
			1,
			1,
			2,
			[][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
		},
	}

	for _, tt := range testCases {
		newImage := floodFill(tt.image, tt.sr, tt.sc, tt.newColor)
		fmt.Println(newImage)
		if !reflect.DeepEqual(newImage, tt.newImage) {
			t.Errorf("error")
		}
	}
}
