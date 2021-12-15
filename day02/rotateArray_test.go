package day02

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	if !reflect.DeepEqual(rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3), []int{5, 6, 7, 1, 2, 3, 4}) {
		t.Errorf("error")
	}
	if !reflect.DeepEqual(rotate([]int{-1, -100, 3, 99}, 2), []int{3, 99, -1, -100}) {
		t.Errorf("error")
	}
	if !reflect.DeepEqual(rotate([]int{-1}, 2), []int{-1}) {
		t.Errorf("error")
	}

}
