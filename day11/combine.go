package day11

import "fmt"

/*
77. 组合
给定两个整数n和k，返回范围[1, n]中所有可能的k个数的组合。
你可以按任何顺序返回答案
*/

func combine(n int, k int) [][]int {
	var res [][]int
	if n < k || n <= 0 || k <= 0 {
		return res
	}

	res = search(make([]int, 0), res, 1, n, k)

	return res
}

func search(tmp []int, res [][]int, index int, n int, k int) [][]int {
	if len(tmp) == k {
		slice1 := make([]int, len(tmp))

		copy(slice1, tmp)

		res = append(res, slice1)
		fmt.Println(res)
		return res
	}

	if len(tmp)+n-index+1 < k {
		return res
	}
	for i := index; i <= n; i++ {
		tmp = append(tmp, i)
		res = search(tmp, res, i+1, n, k)
		tmp = tmp[0 : len(tmp)-1]
	}

	return res
}
