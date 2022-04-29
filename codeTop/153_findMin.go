package codeTop

/*
153. 寻找旋转排序数组中的最小值
已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。例如，原数组 nums = [0,1,2,4,5,6,7] 在变化后可能得到：
若旋转 4 次，则可以得到 [4,5,6,7,0,1,2]
若旋转 7 次，则可以得到 [0,1,2,4,5,6,7]
注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。

给你一个元素值 互不相同 的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。

你必须设计一个时间复杂度为O(log n) 的算法解决此问题。

*/

func findMin(nums []int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2

		// 如果是标准从小到大排列0，1，2，3 或者 3， 0， 1， 2
		if mid == l && nums[l] <= nums[r] || mid > 0 && nums[mid-1] > nums[mid] && nums[mid] < nums[r] {
			return nums[mid]
		}
		// 如果是标准从小到大排列，或者mid左边 < 右指针，r-
		if nums[mid] > nums[l] && nums[mid] < nums[r] || mid > 0 && nums[mid-1] < nums[r] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return nums[l] // 无用
}