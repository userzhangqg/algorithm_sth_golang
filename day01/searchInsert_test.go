package day01

import "testing"

func TestSearchInsert(t *testing.T) {
	if searchInsert([]int{1, 3, 5, 6}, 5) != 2 {
		t.Errorf("error")
	}

	if searchInsert([]int{1, 3, 5, 6}, 2) != 1 {
		t.Errorf("error")
	}
	if searchInsert([]int{1, 3, 5, 6}, 7) != 4 {
		t.Errorf("error")
	}
	if searchInsert([]int{1, 3, 5, 6}, 0) != 0 {
		t.Errorf("error")
	}
	if searchInsert([]int{1}, 0) != 0 {
		t.Errorf("error")
	}
}
