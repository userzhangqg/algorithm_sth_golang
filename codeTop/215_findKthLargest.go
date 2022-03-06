package codeTop

/*
215. 数组中的第k个最大元素

给定整数数组nums和整数k，请返回数组中第k个最大的元素。
请注意，你需要找的是数组排序后的第k个最大的元素，而不是第k个不同的元素。
*/

func findKthLargest(nums []int, k int) int {
	quickSort(nums, 0, len(nums)-1)
	return nums[k-1]
}

func quickSort(array []int, L, R int) {
	if L >= R {
		return
	}

	//以当前数组的最后一个元素(1)作为中枢pivot，进行划分
	left, right, pivot := L, R, array[L]
	for left < right {
		for left < right && array[right] <= pivot {
			right--
		}
		if left < right {
			//将比中枢值大的移动到左端l处 由于l处为中枢或者该位置值已经被替换到r处，所以直接可以替换
			array[left] = array[right]
		}

		//fmt.Println(array)

		for left < right && array[left] >= pivot {
			left++
		}
		if left < right {
			//将比中枢值小的移动到右端r处 由于前面r处的值已经换到l处，所以该位置值也可以替换掉
			array[right] = array[left]
		}

		//fmt.Println(array)
	}

	//l==r时，重合，这个位置就是中枢的最终位置
	array[left] = pivot

	quickSort(array, L, left-1)
	quickSort(array, left+1, R)
}
