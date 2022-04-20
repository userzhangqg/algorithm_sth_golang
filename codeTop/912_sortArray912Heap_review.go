package codeTop

func sortArray912HeapReview(nums []int) []int {
	end := len(nums) - 1

	for i := end / 2; i >= 0; i-- {
		heapSortReview(nums, i, end)
	}

	for i := end; i >= 0; i-- {
		nums[0], nums[end] = nums[end], nums[0]
		end--
		heapSortReview(nums, 0, end)
	}
	return nums
}

func heapSortReview(nums []int, root int, end int) {
	for {
		child := root*2 + 1

		if child > end {
			return
		}

		if child < end && nums[child] < nums[child+1] {
			child++
		}

		if nums[child] > nums[root] {
			nums[child], nums[root] = nums[root], nums[child]
		}
		root = child
	}
}
