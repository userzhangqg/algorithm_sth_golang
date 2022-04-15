package codeTop

func longestConsecutiveReview(nums []int) int {
	hashMap := make(map[int]bool)

	for _, n := range nums {
		hashMap[n] = true
	}

	var maxSize int

	for k, _ := range hashMap {

		if !hashMap[k-1] {
			curIndex := k
			curSize := 1

			for hashMap[curIndex+1] {
				curIndex++
				curSize++
			}

			if curSize > maxSize {
				maxSize = curSize
			}
		}
	}
	return maxSize
}
