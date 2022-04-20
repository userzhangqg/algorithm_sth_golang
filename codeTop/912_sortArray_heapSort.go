package codeTop

func sortArray912Heap(nums []int) []int {
	end := len(nums) - 1
	for i := end / 2; i >= 0; i-- {
		heapSort(nums, i, end)
	}

	for i := end; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		end--
		heapSort(nums, 0, end)
	}
	return nums
}

func heapSort(nums []int, root int, end int) {
	for {
		child := root*2 + 1

		if child > end {
			return
		}

		if child < end && nums[child] < nums[child+1] {
			child++
		}
		if nums[root] > nums[child] {
			return
		}
		nums[root], nums[child] = nums[child], nums[root]
		root = child
	}
}
