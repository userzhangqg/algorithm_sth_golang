package codeTop

import "sort"

/*
https://leetcode-cn.com/problems/3sum/
15. 三数之和

给你一个包含n个整数的数组nums，判断nums中是否存在三个元素a，b，c，使得a+b+c = 0 ？
请你找出所有和为0且不重复的三元组。
注意：答案中不可以包含重复的三元组。
*/

/*
自我思路：先排序 + ？

正确思路：排序 + 双指针
*/

func threeSum(nums []int) [][]int {
	var ans [][]int

	if nums == nil || len(nums) < 3 {
		return ans
	}

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		L := i + 1
		R := len(nums) - 1

		// 如果当前数字大于0，则三数之和一定大于0，所以结束循环
		if nums[i] > 0 {
			break
		}

		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for L < R {
			if nums[i]+nums[L]+nums[R] == 0 {
				ans = append(ans, []int{nums[i], nums[L], nums[R]})
				// 去重
				for L < R && nums[L] == nums[L+1] {
					L++
				}
				// 去重
				for L < R && nums[R] == nums[R-1] {
					R--
				}
				L++
				R--
			} else if nums[i]+nums[L]+nums[R] < 0 {
				L++
			} else {
				R--
			}
		}
	}

	return ans
}
