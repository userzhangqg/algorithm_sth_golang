package codeTop

func majorityElementReview(nums []int) int {
	var major int
	var count int

	for _, v := range nums {
		if count == 0 {
			major = v
		}

		if major == v {
			count++
		} else {
			count--
		}
	}
	return major
}
