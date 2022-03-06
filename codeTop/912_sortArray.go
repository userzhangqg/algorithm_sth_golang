package codeTop

/*
912. 排序数组
给你一个整数数组nums，请你将该数组升序排列
https://leetcode-cn.com/problems/sort-an-array/

十大排序算法-golang实现
https://leetcode-cn.com/problems/sort-an-array/solution/shi-da-pai-xu-suan-fa-golangshi-xian-by-8p1bl/
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

		//if l < r {
		//	nums[r] = nums[l]
		//}

		for l < r && nums[r] >= p {
			r--
		}
		//if l < r {
		//	nums[l] = nums[r]
		//}

		nums[l], nums[r] = nums[r], nums[l]
	}

	//nums[r] = p
	nums[r], nums[R] = nums[R], nums[r]
	quickSort912(nums, L, r-1)
	quickSort912(nums, r+1, R)
}
