package codeTopReview

func merge(nums1 []int, m int, nums2 []int, n int) {
	var sorted []int

	i1, i2 := 0, 0

	for {
		if i1 >= m {
			sorted = append(sorted, nums2[i2:]...)
			break
		}
		if i2 >= n {
			sorted = append(sorted, nums1[i1:]...)
			break
		}

		if nums1[i1] < nums2[i2] {
			sorted = append(sorted, nums1[i1])
			i1++
		} else {
			sorted = append(sorted, nums2[i2])
			i2++
		}
	}
	copy(nums1, sorted)
}
