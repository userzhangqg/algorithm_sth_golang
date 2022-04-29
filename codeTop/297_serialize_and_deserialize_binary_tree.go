package codeTop

import (
	"strconv"
	"strings"
)

/*
297. 二叉树的序列化与反序列化

序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将装换后的数据存储在一个文件或内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。

请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列、反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

提示：输入输出格式与LeetCode目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。
*/

type Codec struct {
}

func Constructor297() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	builder := strings.Builder{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			builder.WriteString("null,")
			return
		}
		builder.WriteString(strconv.Itoa(root.Val))
		builder.WriteByte(',')
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return builder.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	s := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if s[0] == "null" {
			s = s[1:]
			return nil
		}
		val, _ := strconv.Atoi(s[0])
		s = s[1:]
		return &TreeNode{val, build(), build()}
	}
	return build()
}
