package codeTopReview

/*
704. 二分查找
https://leetcode.cn/problems/binary-search/
*/

func search704(nums []int, target int) int {
	l, r := 0, len(nums)-1

	// 注意一个元素的情况
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
