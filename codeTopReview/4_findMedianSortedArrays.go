package codeTopReview

/*
4. 寻找两个正序数组的中位数
https://leetcode.cn/problems/median-of-two-sorted-arrays/

时间复杂度：遍历全部数组
O(m+n)

空间复杂度：开辟了一个数组，保存合并后的两个数组
O(m+n)
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := 0, 0
	var ans []int

	for l1 < len(nums1) && l2 < len(nums2) {
		if nums1[l1] < nums2[l2] {
			ans = append(ans, nums1[l1])
			l1++
		} else {
			ans = append(ans, nums2[l2])
			l2++
		}
	}

	if l1 < len(nums1) {
		ans = append(ans, nums1[l1:]...)
	} else if l2 < len(nums2) {
		ans = append(ans, nums2[l2:]...)
	}

	if len(ans)%2 == 0 {
		return float64(ans[len(ans)/2-1]+ans[len(ans)/2]) / 2
	} else {
		return float64(ans[len(ans)/2])
	}
}
