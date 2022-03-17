package codeTop

func trapReview(height []int) int {
	l, r := 0, len(height)-1
	l_max, r_max, ans := 0, 0, 0

	for l < r {
		if height[l] < height[r] {
			if height[l] >= l_max {
				l_max = height[l]
			} else {
				ans = l_max - height[l] + ans
			}
			l++
		} else {
			if height[r] >= r_max {
				r_max = height[r]
			} else {
				ans = r_max - height[r] + ans
			}
			r--
		}
	}
	return ans
}
