package codeTop

/*
!!! TODO
快排忘记
堆排序学习
*/

func sortArray912(nums []int) []int {
	quickSort912Review0419(nums, 0, len(nums)-1)
	return nums
}

func quickSort912Review0419(nums []int, L, R int) {
	if L >= R {
		return
	}

	l, r, p := L, R, nums[R]

	for l < r {
		// 大数放右边
		for l < r && nums[l] <= p {
			l++
		}
		if l < r {
			nums[r] = nums[l]
		}

		// 小数放左边
		for l < r && nums[r] >= p {
			r--
		}
		if l < r {
			nums[l] = nums[r]
		}
	}

	nums[l] = p
	// 递归排序左右两边
	quickSort912Review0419(nums, L, l-1)
	quickSort912Review0419(nums, l+1, R)
}
