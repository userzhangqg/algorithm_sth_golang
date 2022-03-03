package codeTop

/*
215. 数组中的第k个最大元素

给定整数数组nums和整数k，请返回数组中第k个最大的元素。
请注意，你需要找的是数组排序后的第k个最大的元素，而不是第k个不同的元素。
*/

func findKthLargest(nums []int, k int) int {
	quickSort(nums, 0, len(nums)-1)
	return nums[k-1]
}

func quickSort(array []int, L, R int) {
	if L >= R {
		return
	}
	left, right, pivot := L, R, array[L]
	for left < right {
		for left < right && array[right] <= pivot {
			right--
		}
		if left < right {
			array[left] = array[right]
		}

		for left < right && array[left] >= pivot {
			left++
		}
		if left < right {
			array[right] = array[left]
		}

		if left >= right {
			array[left] = pivot
		}
	}
	quickSort(array, L, right-1)
	quickSort(array, right+1, R)
}
