package codeTop

/*
88. 合并两个有序数组
https://leetcode-cn.com/problems/merge-sorted-array/submissions/
给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。

请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。

注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		//nums1 = nums2
		copy(nums1, nums2)
	}
	if n == 0 {
		return
	}
	nums1 = append(nums1[:m], nums2...)

	L, R := 0, m+n-1

	quickSort88(nums1, L, R)

}

func quickSort88(nums []int, L int, R int) {
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
	quickSort88(nums, L, r-1)
	quickSort88(nums, r+1, R)
}
