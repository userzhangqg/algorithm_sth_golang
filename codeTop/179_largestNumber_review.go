package codeTop

import (
	"sort"
	"strconv"
)

func largestNumberReview(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := strconv.Itoa(nums[i]), strconv.Itoa(nums[j])
		str1, str2 := x+y, y+x
		for k := 0; k < len(str1); k++ {
			if str1[k] == str2[k] {
				continue
			}
			return str1[k] > str2[k]
		}
		return true
	})
	if nums[0] == 0 {
		return "0"
	}
	var ans string
	for _, v := range nums {
		ans += strconv.Itoa(v)
	}
	return ans
}
