package codeTop

func mySqrtReview(x int) int {
	l, r := 0, x
	var ans int

	for l <= r {
		mid := l + (r-l)/2

		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}
