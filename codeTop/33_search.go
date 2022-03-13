package codeTop

/*
33. 搜索旋转排序数组
https://leetcode-cn.com/problems/search-in-rotated-sorted-array/

整数数组nums按升序排列，数组中的值 互不相同。
...
给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。

!!! 做过忘记

思路：
1. 二分法 注意判别边界
*/

func search33(nums []int, target int) int {
	l, r := 0, len(nums)-1

	// 注意边界，两指针相撞
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			// 注意特性，中间值比左指针大时，目标值一定在右边
			if nums[mid] > nums[l] || (nums[r] > nums[mid] && nums[r] >= target) {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			// 同上，中间值比右指针小时，目标值一定在左边
			if nums[mid] < nums[r] || (nums[l] < nums[mid] && nums[l] <= target) {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}
