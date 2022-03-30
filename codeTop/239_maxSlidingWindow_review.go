package codeTop

func maxSlidingWindowReview(nums []int, k int) []int {
	size := len(nums)
	var ans []int
	if size == 0 || k < 0 || k > size {
		return ans
	}

	var q []int
	for right := 0; right < size; right++ {
		for len(q) > 0 && nums[right] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}

		q = append(q, right)

		left := right - k + 1

		// 如果窗口左边界超过左侧边框，切除
		if q[0] < left {
			q = q[1:]
		}

		// 如果窗口达到k进行数据添加
		if right+1 >= k {
			ans = append(ans, nums[q[0]])
		}
	}
	return ans
}
