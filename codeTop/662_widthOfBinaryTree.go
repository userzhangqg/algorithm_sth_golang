package codeTop

/*
662. 二叉树最大宽度

给定一个二叉树，编写一个函数来获取这个树的最大宽度。
树的宽度是所有层中的最大宽度。这个二叉树与满二叉树结构相同，但一些节点为空。
每一层的宽度被定义为两个端点（该层最左和最右的非空节点，两端点间的null节点也计入长度）之间的长度。

*/

func widthOfBinaryTree(root *TreeNode) int {
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
		if pos-hashMap[depth]+1 > ans {
			ans = pos - hashMap[depth] + 1
		}
		dfs(root.Left, depth+1, pos*2)
		dfs(root.Right, depth+1, pos*2+1)
	}
	dfs(root, 0, 0)
	return ans
}
