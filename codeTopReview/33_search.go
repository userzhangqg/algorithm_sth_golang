package codeTopReview

/*
33. 搜索旋转排序数组
https://leetcode.cn/problems/search-in-rotated-sorted-array/

时间复杂度： O(\log n)O(logn)，其中 nn 为 \textit{nums}nums 数组的大小。整个算法时间复杂度即为二分查找的时间复杂度 O(\log n)O(logn)。

空间复杂度： O(1)O(1) 。我们只需要常数级别的空间存放变量。

*/

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			if nums[mid] > nums[l] || (nums[r] > nums[mid] && nums[r] >= target) {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			if nums[mid] < nums[r] || (nums[l] < nums[mid] && nums[l] <= target) {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}
