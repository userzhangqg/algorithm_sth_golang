package codeTop

import (
	"strconv"
	"strings"
)

type CodecReview struct {
}

func ConstructorReview297() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize297(root *TreeNode) string {
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
func (this *Codec) deserialize297(data string) *TreeNode {
	s := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if s[0] == "null" {
			s = s[1:]
			return nil
		}
		n, _ := strconv.Atoi(s[0])
		s = s[1:]
		return &TreeNode{n, build(), build()}
	}
	return build()
}
