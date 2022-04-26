package codeTop

func widthOfBinaryTreeReview(root *TreeNode) int {
	var ans int
	hashMap := make(map[int]int)

	var dfs func(root *TreeNode, depth int, pos int)
	dfs = func(root *TreeNode, depth int, pos int) {
		if root == nil {
			return
		}
		if _, ok := hashMap[depth]; !ok {
			hashMap[depth] = pos
		}
		dfs(root.Left, depth+1, pos*2)
		dfs(root.Right, depth+1, pos*2+1)
		if pos-hashMap[depth]+1 > ans {
			ans = pos - hashMap[depth] + 1
		}
	}
	dfs(root, 0, 0)
	return ans
}
