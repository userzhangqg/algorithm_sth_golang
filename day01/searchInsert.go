package day01

/*
35. 搜索插入的位置
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将被按顺序插入的位置。
请必须使用时间复杂度为O(log n) 的算法
*/

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = l + 1
		}
	}
	return l
}
