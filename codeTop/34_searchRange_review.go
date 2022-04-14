package codeTop

func searchRangeReview(nums []int, target int) []int {
	L, R := 0, len(nums)-1

	var l, r int
	lIndex, RIndex := -1, -1

	for l, r = L, R; l <= r; {
		mid := l + (r-l)/2
		if nums[mid] == target {
			lIndex = mid
			r = mid - 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	// 注意小于等于，l, r 从头开始
	for l, r = L, R; l <= r; {
		mid := l + (r-l)/2
		if nums[mid] == target {
			RIndex = mid
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return []int{lIndex, RIndex}
}
