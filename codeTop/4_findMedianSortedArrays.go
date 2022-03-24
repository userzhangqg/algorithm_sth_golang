package codeTop

/*
4. 寻找出两个正序数组的中位数
https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
给定两个大小分别为m和n的正序（从小到大）数组nums1 和 nums2。请你找出并返回这两个正序数组中的中位数。
算法时间复杂度应该为O（log(m+n)）

！！！ 待学习
思路：
归并排序取中

正确思路：
归并排序复杂度不符合要求，使用二分法，待学习
https://leetcode-cn.com/problems/median-of-two-sorted-arrays/solution/gui-bing-pai-xu-shuang-zhi-zhen-er-fen-c-4dwb/
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var ans []int

	l1, l2 := 0, 0

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
		return float64(ans[len(ans)/2-1]+ans[len(ans)/2]) / 2.0
	} else {
		return float64(ans[len(ans)/2])
	}
}

func FindMedianSortedArraysDemo(nums1, nums2 []int) float64 {
	var res []int
	m, n := len(nums1), len(nums2)
	l1, l2 := 0, 0
	for l1 < m && l2 < n {
		if nums1[l1] < nums2[l2] {
			res = append(res, nums1[l1])
			l1++
		} else {
			res = append(res, nums2[l2])
			l2++
		}
	}
	res = append(res, nums1[l1:]...)
	res = append(res, nums2[l2:]...)
	length := m + n
	if length%2 == 1 {
		return float64(res[length/2])
	}
	mid1 := res[length/2]
	mid2 := res[length/2-1]
	return float64(mid1+mid2) / 2.0
}
