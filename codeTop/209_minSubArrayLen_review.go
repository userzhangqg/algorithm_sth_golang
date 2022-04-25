package codeTop

func minSubArrayLenReview(target int, nums []int) int {
	var sum int
	l, ans := 0, len(nums)+1

	for i, v := range nums {
		if v == target {
			return 1
		}
		sum += v
		for sum >= target && sum-nums[l] >= target {
			sum -= nums[l]
			l++
		}
		if sum >= target && i-l+1 < ans {
			ans = i - l + 1
		}
	}
	if ans == len(nums)+1 {
		return 0
	}
	return ans
}
