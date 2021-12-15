package day02

import (
	"reflect"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	res := sortedSquares([]int{-4, -1, 0, 3, 10})
	if !reflect.DeepEqual(res, []int{0, 1, 9, 16, 100}) {
		t.Errorf("error")
	}
	//if sortedSquares([]int{-4, -1, 0, 3, 10}) != []int{0, 1, 9, 16, 100} {
	//	t.Errorf("error")
	//}
}
