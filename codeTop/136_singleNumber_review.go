package codeTop

func singleNumberReview(nums []int) int {
	var ans int

	for _, v := range nums {
		ans = ans ^ v
	}
	return ans
}
