package codeTop

/*
34. 在排序数组中查找元素的第一个和最后一个位置

给定一个按照升序排列的整数数组nums，和一个目标值target。找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值target，返回[-1, -1]

进阶：你可以设计并实现时间复杂度为O(log n)的算法解决此问题了吗

!!! TODO
思路：
1. 二分查找后左右扩张（做过，细节错误）
*/

func searchRange(nums []int, target int) []int {
	for l, r := 0, len(nums)-1; l <= r; {
		mid := l + (r-l)/2
		if nums[mid] == target {
			l = mid
			r = mid
			for l > 0 && nums[l-1] == target {
				l--
			}
			for r < len(nums)-1 && nums[r+1] == target {
				r++
			}
			return []int{l, r}
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return []int{-1, -1}
}
