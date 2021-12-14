package day01

import "testing"

func TestSearch(t *testing.T) {
	//nums := []int{-1, 0, 3, 5, 9, 12}
	//target := 9
	if search([]int{-1, 0, 3, 5, 9, 12}, 9) != 4 {
		t.Errorf("error")
	}

	if search1([]int{-1, 0, 3, 5, 9, 12}, 9) != 4 {
		t.Errorf("error")
	}

	if search1([]int{-1, 0, 3, 5, 9, 12}, 2) != -1 {
		t.Errorf("error")
	}
}
