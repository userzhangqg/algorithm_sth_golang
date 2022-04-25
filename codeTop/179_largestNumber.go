package codeTop

import (
	"sort"
	"strconv"
)

/*
179. 最大数

给定一组非负整数nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数

注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数

!!! TODO
思路：
1. 排序后遍历
https://leetcode-cn.com/problems/largest-number/solution/si-lu-qing-xi-dai-ma-jian-dan-by-damon-s-4ufd/
*/

func largestNumber(nums []int) string {
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
