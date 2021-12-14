package day01

import "testing"

func TestFirstBadVersion(t *testing.T) {
	if firstBadVersion(5) != 4 {
		t.Errorf("err")
	}
}
