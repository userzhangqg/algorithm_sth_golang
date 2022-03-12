package codeTop

func merge88(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
	}
	if n == 0 {
		return
	}

	nums1 = append(nums1[:m], nums2...)
	l, r := 0, len(nums1)-1
	quickSortReview(nums1, l, r)
}

func quickSortReview(nums []int, L, R int) {
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
	quickSortReview(nums, L, r-1)
	quickSortReview(nums, r+1, R)
}
