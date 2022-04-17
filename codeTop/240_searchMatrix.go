package codeTop

/*
240. 搜索二维矩阵II
https://leetcode-cn.com/problems/search-a-2d-matrix-ii/
编写一个高效的算法来搜索m x n矩阵matrix中的一个目标值target。
该矩阵具有以下特性：
- 每行的元素从左到右升序排列
- 每列的元素从上到下升序排列

!!! TODO
思路错误。非正方形矩阵
思路：
1. 每行或每列二分查找
2. Z字形查找
*/

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])

	x, y := 0, n-1

	for x < m && y >= 0 {
		if matrix[x][y] == target {
			return true
		}

		if matrix[x][y] > target {
			// 先确定纵坐标
			y--
		} else {
			// 再确定横坐标
			x++
		}
	}
	return false
}

/*

// 审题错误，非正方形矩阵
func searchMatrix(matrix [][]int, target int) bool {
	size := len(matrix)

	for i := 0; i < size; i++ {
		if matrix[i][i] == target {
			return true
		}else if i > 0 && target > matrix[i-1][i-1] && target < matrix[i][i] {
			return binarySearch240(matrix[i][:i], target) || binarySearchC240(matrix, i, target)
		}
	}
	return false
}

func binarySearch240(nums []int, target int) bool {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r - l) / 2
		if nums[mid] == target {
			return true
		}else if nums[mid] < target {
			l = mid + 1
		}else {
			r = mid - 1
		}
	}
	return false
}

func binarySearchC240(matrix [][]int, p int, target int) bool {
	l, r := 0, p - 1

	for l <= r {
		mid := l + (r - l) / 2
		if matrix[mid][p] == target {
			return true
		}else if matrix[mid][p] < target {
			l = mid + 1
		}else {
			r = mid - 1
		}
	}
	return false
}
*/
