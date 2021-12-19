package day03

/*
283. 移动零

给定一个数组nums，编写一个函数将所有的0移动到数组的末尾， 同时保持非零元素的相对顺序
*/

func moveZeroes(nums []int) {
	if len(nums) < 2 {
		return
	}
	//l, r, k := 0, len(nums) - 1, 1
	for l, r, k := 0, len(nums)-1, 1; k <= r; k = k + 1 {
		if nums[l] == 0 {
			if nums[k] != 0 {
				nums[l], nums[k] = nums[k], nums[l]
				l++
			}
		} else {
			l++
		}
	}
}

func moveZeroesDemo(nums []int) {
	/*
		官解：
		https://leetcode-cn.com/problems/move-zeroes/solution/yi-dong-ling-by-leetcode-solution/
		使用双指针，左指针指向当前已经处理好的序列的尾部，右指针指向待处理序列的头部。

		右指针不断向右移动，每次右指针指向非零数，则将左右指针对应的数交换，同时左指针右移。

		注意到以下性质：

		左指针左边均为非零数；

		右指针左边直到左指针处均为零。

		因此每次交换，都是将左指针的零与右指针的非零数交换，且非零数的相对顺序并未改变。
	*/
	//left, right, n := 0, 0, len(nums)
	for left, right, n := 0, 0, len(nums); right < n; right = right + 1 {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		//right++
	}
}
