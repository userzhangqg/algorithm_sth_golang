package codeTopReview

/*
215. 数组中的第K个最大元素
https://leetcode.cn/problems/kth-largest-element-in-an-array/
*/

func findKthLargest(nums []int, k int) int {
	quickSort(nums, 0, len(nums)-1)
	return nums[k-1]
}

func quickSort(nums []int, L int, R int) {

	if L >= R {
		return
	}
	p, left, right := nums[R], L, R
	for left < right {
		//for left < right && nums[right] <= p {
		//	right--
		//}
		//if left < right {
		//	nums[left] = nums[right]
		//}
		//
		//for left < right && nums[left] >= p {
		//	left++
		//}
		//if left < right {
		//	nums[right] = nums[left]
		//}
		for left < right && nums[left] >= p {
			left++
		}
		if left < right {
			nums[right] = nums[left]
		}

		for left < right && nums[right] <= p {
			right--
		}
		if left < right {
			nums[left] = nums[right]
		}

	}

	nums[left] = p

	quickSort(nums, L, left-1)
	quickSort(nums, left+1, R)
}
