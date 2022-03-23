package codeTop

/*
给你两个单词word1 和word2， 请返回将word1转换成word2 所使用的最少操作数 。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/edit-distance

思路：
动态规划
https://leetcode-cn.com/problems/edit-distance/solution/edit-distance-by-ikaruga/
https://leetcode-cn.com/problems/edit-distance/solution/pythongo-dong-tai-gui-hua-by-himymben-3nce/
https://leetcode-cn.com/problems/edit-distance/solution/dai-ma-sui-xiang-lu-72-bian-ji-ju-chi-do-y87e/
*/

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for d := range dp {
		dp[d] = make([]int, n+1)
	}

	for i := 0; i < m+1; i++ {
		dp[i][0] = i
	}
	for j := 0; j < n+1; j++ {
		dp[0][j] = j
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			// 注意i, j 要减一
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 本步数为上一最小步数+1
				dp[i][j] = min72(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[m][n]
}

func min72(args ...int) int {
	ans := args[0]

	for _, v := range args {
		if v < ans {
			ans = v
		}
	}
	return ans
}
