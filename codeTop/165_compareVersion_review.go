package codeTop

import (
	"strconv"
	"strings"
)

func compareVersionReview(version1 string, version2 string) int {
	arr1 := strings.Split(version1, ".")
	arr2 := strings.Split(version2, ".")

	size1 := len(arr1)
	size2 := len(arr2)

	for i, j := 0, 0; i < size1 || j < size2; i, j = i+1, j+1 {
		x, y := 0, 0
		if i < size1 {
			x, _ = strconv.Atoi(arr1[i])
		}
		if j < size2 {
			y, _ = strconv.Atoi(arr2[j])
		}

		if x > y {
			return 1
		} else if x < y {
			return -1
		}
	}

	return 0
}
