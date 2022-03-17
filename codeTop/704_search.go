package codeTop

/*
704. 二分查找
https://leetcode-cn.com/problems/binary-search/

给定一个n个元素有序的升序整数数组nums和一个目标值target，写一个函数搜索nums中target，如果目标值存在返回下标，否则返回-1
*/

func search704(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
