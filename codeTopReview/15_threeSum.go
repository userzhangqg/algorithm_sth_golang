package codeTopReview

import "sort"

/*
15. 三数之和
https://leetcode.cn/problems/3sum/

思路：
1. 排序 + 双指针
*/

func threeSum(nums []int) [][]int {
	var ans [][]int

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1

		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for left < right {
			if nums[i]+nums[left]+nums[right] == 0 {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return ans
}
