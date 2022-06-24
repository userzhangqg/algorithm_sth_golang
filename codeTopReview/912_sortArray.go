package codeTopReview

/*
912. 排序数组
https://leetcode.cn/problems/sort-an-array/
*/

func sortArray(nums []int) []int {
	quickSort912(nums, 0, len(nums)-1)
	return nums
}

func quickSort912(nums []int, L, R int) {
	if L >= R {
		return
	}

	l, r, p := L, R, nums[R]

	for l < r {
		for l < r && nums[l] <= p {
			l++
		}
		if l < r {
			nums[r] = nums[l]
		}

		for l < r && nums[r] >= p {
			r--
		}
		if l < r {
			nums[l] = nums[r]
		}
	}
	nums[l] = p
	quickSort912(nums, L, l-1)
	quickSort912(nums, l+1, R)
}
